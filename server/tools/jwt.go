package tools

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaim struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return "aSecret"
	}

	return secret
}

func GenerateJwt(ctx context.Context, userID int, email string) (string, error) {
	// create a new token with custom claims that is passed as a second parameter
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID:    userID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	// after the token has been created, sign it with JWT_SECRET to make it secure
	token, err := tokenClaims.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return token, nil
}
