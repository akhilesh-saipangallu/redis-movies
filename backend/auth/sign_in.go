package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleSignIn(c *gin.Context) {
	var (
		err         error
		requestData SignInRequest
	)
	err = c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	validate := validator.New()
	if err = validate.Struct(requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	user, err := getUserDetails(c.Request.Context(), requestData.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid credentials"})
		return
	}

	// verify password
	if !verifyPassword(user.Password, requestData.Password) {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid credentials"})
		return
	}

	// generate token
	jwtToken, err := generateJWT(*user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
