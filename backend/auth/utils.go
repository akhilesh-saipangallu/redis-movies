package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func verifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func generateJWT(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		err = fmt.Errorf("generateJWT: failed to get token: %w", err)
		return "", err
	}
	return tokenString, err
}

func GetCurrentUserId(c *gin.Context) (string, bool) {
	userData, ok := c.Get("user_data")
	if !ok {
		return "", ok
	}

	return userData.(UserPartial).Id, ok
}
