package service

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

func (s *Service) Validate(token string) (models.TokenInfo, error) {
	token_, err := s.ParseToken(token)
	if err != nil {
		return models.TokenInfo{}, err
	}

	return token_, nil
}

func (s *Service) Refresh(token string) (string, error) {
	if err := s.repo.Refresh(token); err != nil {
		return "", err
	}

	tokenInfo, err := s.ParseToken(token)
	if err != nil {
		return "", err
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:  tokenInfo.UserId,
		IsAdmin: tokenInfo.IsAdmin,
	})

	newTokenString, err := newToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return newTokenString, nil
}

func (s *Service) ParseToken(token string) (models.TokenInfo, error) {
	token_, err := jwt.ParseWithClaims(token, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return models.TokenInfo{}, err
	}

	claims, ok := token_.Claims.(*models.TokenClaims)
	if !ok {
		return models.TokenInfo{}, errors.New("token claims are not of type tokenClaims")
	}

	return models.TokenInfo{UserId: claims.UserId, IsAdmin: claims.IsAdmin}, nil
}

func (s *Service) IsTokenInvalid(token string) (bool, error) {
	return s.repo.IsTokenInvalid(token)
}
