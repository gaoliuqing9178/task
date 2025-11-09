package jwtUtil

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(id string, expireTime time.Duration, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(expireTime).Unix(),
	})

	signedString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func ParseToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), nil
	}

	return "", jwt.ErrInvalidKey
}
