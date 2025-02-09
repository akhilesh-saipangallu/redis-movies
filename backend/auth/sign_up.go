package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()
	var (
		err         error
		requestData SignUpRequest
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

	// check if user already exists
	userExists, err := doesUserExists(ctx, requestData.Email)

	if err != nil {
		log.Printf("HandleSignUp: %v", err)
		http.Error(w, newAuthErrorResponseJson("internal error"), http.StatusInternalServerError)
		return
	}
	if userExists {
		http.Error(w, newAuthErrorResponseJson("user already exists"), http.StatusBadRequest)
		return
	}

	// create the user record in DB
	createUser(ctx, requestData)
	w.WriteHeader(http.StatusCreated)
}
