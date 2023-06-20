package api

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"simpleProject/pkg/model"
	"strings"
	"time"
)

func (s *service) CheckAuthLogin(data *model.LoginRequest) (*model.Auth, error) {
	var res model.Auth

	query := `SELECT author_id, name, role, login, pswd_hash FROM authors WHERE login = $1`
	err := s.store.repo.Get(&res, query, false, data.Login)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, fmt.Errorf("authentication failed")
	}

	if (data.Login == res.Login) && (data.Password == res.PswdHash) {
		return &res, nil
	}

	return nil, fmt.Errorf("authentication failed")
}

func (s *service) CheckAuthRefresh(data jwt.MapClaims) (*model.Auth, error) {
	var res model.Auth

	ID := data["sub"].(string)
	login := data["login"].(string)

	query := `SELECT author_id, name, role, login, pswd_hash FROM authors WHERE author_id = $1`
	err := s.store.repo.Get(&res, query, false, ID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, fmt.Errorf("authentication failed")
	}

	if (ID == res.ID) && (login == res.Login) {
		return &res, nil
	}

	return nil, fmt.Errorf("authentication failed")
}

func (s *service) GenerateJWT(data *model.Auth) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   data.ID,
		"name":  data.Name,
		"role":  data.Role,
		"login": data.Login,
		"exp":   time.Now().Add(time.Minute * 55).Unix(),
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

	if refresh {
		return token, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *service) ParseJWT(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("error JWT parse")
}
