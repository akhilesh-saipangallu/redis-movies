package auth

import (
	"context"
	"fmt"
	"time"

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

func verifyJWT(ctx context.Context, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		return fmt.Errorf("verifyJWT: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		ctx = context.WithValue(ctx, "user", UserPartial{
			Id:    claims["id"].(string),
			Email: claims["email"].(string),
		})
	} else {
		return fmt.Errorf("verifyJWT: bad claims")
	}
	return nil
}
