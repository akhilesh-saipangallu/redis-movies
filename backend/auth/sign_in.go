package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()
	var (
		err         error
		requestData SignInRequest
	)
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, newAuthErrorResponseJson("bad request"), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err = validate.Struct(requestData); err != nil {
		http.Error(w, newAuthErrorResponseJson("bad request"), http.StatusBadRequest)
		return
	}

	user, err := getUserDetails(ctx, requestData.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, newAuthErrorResponseJson("invalid credentials"), http.StatusForbidden)
		return
	}

	// verify password
	if !verifyPassword(user.Password, requestData.Password) {
		http.Error(w, newAuthErrorResponseJson("invalid credentials"), http.StatusForbidden)
		return
	}

	// generate token
	jwtToken, err := generateJWT(*user)
	if err != nil {
		log.Println(err)
		http.Error(w, newAuthErrorResponseJson("invalid credentials"), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(newSignInResponseJson(jwtToken)))
}
