package auth

import "encoding/json"

type SignUpRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type UserPartial struct {
	Id    string
	Email string
}

type AuthErrorResponse struct {
	ErrorMessage string `json:"error"`
}

func newAuthErrorResponseJson(errorMessage string) string {
	authError := AuthErrorResponse{ErrorMessage: errorMessage}
	authErrorJson, _ := json.Marshal(authError)
	return string(authErrorJson)
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

func newSignInResponseJson(token string) string {
	response := SignInResponse{Token: token}
	responseJson, _ := json.Marshal(response)
	return string(responseJson)
}
