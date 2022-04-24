package tools

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaim struct {
	UserID  int    `json:"userId"`
	Email   string `json:"email"`
	TokenID string `json:"uuid"`
	jwt.StandardClaims
}

var accessTokenSecret = []byte(getAccessTokenSecret())
var refreshTokenSecret = []byte(getRefreshTokenSecret())

func getAccessTokenSecret() string {
	secret := os.Getenv("ACCESS_SECRET")

	if secret == "" {
		return ""
	}

	return secret
}

func getRefreshTokenSecret() string {
	secret := os.Getenv("REFRESH_SECRET")

	if secret == "" {
		return ""
	}

	return secret
}

// GenerateAccessToken generates a signed token
func GenerateAccessToken(ctx context.Context, userID int, email string) (string, error) {
	tokenId := uuid.New().String()
	// create a new token with custom claims that is passed as a second parameter
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		UserID:  userID,
		Email:   email,
		TokenID: tokenId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	// after the token has been created, sign it with JWT_SECRET to make it secure
	token, err := tokenClaims.SignedString(refreshTokenSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRefreshToken(ctx context.Context, userID int, email string) (string, error) {
	tokenId := uuid.New().String()
	// create a new token with custom claims that is passed as a second parameter
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		UserID:  userID,
		Email:   email,
		TokenID: tokenId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	// after the token has been created, sign it with JWT_SECRET to make it secure
	token, err := tokenClaims.SignedString(accessTokenSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateJwt validates the token and returns it
func ValidateJwt(ctx context.Context, token string) (*jwt.Token, error) {
	if token == "" {
		return nil, fmt.Errorf("auth token string empty")
	}

	// validate and return the token as *jwt.Token
	// first check the token method
	// check if the signing method is HMAC since secret HS256 to sign the token is used
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there has been a problem with the signing method")
		}

		return accessTokenSecret, nil
	})
}

//func ValidateRefreshToken(ctx context.Context, refreshToken string) (model.UserAuth, error) {
//	token, err := jwt.ParseWithClaims(refreshToken, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
//		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("problem with the signing method %v", t.Header["alg"])
//		}
//
//		return jwtSecret, nil
//	})
//
//	if err != nil {
//		return
//	}
//}
