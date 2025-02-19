package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleSignUp(c *gin.Context) {
	var (
		err         error
		requestData SignUpRequest
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

	// check if user already exists
	userExists, err := doesUserExists(c.Request.Context(), requestData.Email)

	if err != nil {
		log.Printf("HandleSignUp: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	// create the user record in DB
	err = createUser(c.Request.Context(), requestData)
	if err != nil {
		log.Println("error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal error"})
		return
	}

	c.Status(http.StatusCreated)
}
