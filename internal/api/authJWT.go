package api

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

func (s *service) GenerateJWT(user *string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 55).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.cfg.AuthConfig.SecretKeyJWT))
	if err != nil {
		return "", fmt.Errorf("JWT token signing: %w", err)
	}

	return tokenString, nil
}

func (s *service) VerifyJWT(tokenString string, refresh bool) (*jwt.Token, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("token is empty")
	}

	tokenString = strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.cfg.AuthConfig.SecretKeyJWT), nil
	})
	if err != nil {
		return nil, err
	}

	if refresh {
		return token, nil
	}

	return nil, nil
}

func (s *service) ParseJWT(token *jwt.Token) (string, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["sub"].(string), nil
	}
	return "", nil
}
