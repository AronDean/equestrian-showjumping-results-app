package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func JwtDecode(secret, tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Add the user ID to the request context
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)
	return int(userID), nil
}

func JwtEncode(secret string, userID int, expiry time.Time) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expiry.Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
