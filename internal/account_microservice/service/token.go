package service

import (
	"errors"
	"log"
	"os"

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

func (s *Service) Refresh(token string) error {
	return s.repo.Refresh(token)
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

	log.Println(claims)

	return models.TokenInfo{UserId: claims.UserId, IsAdmin: claims.IsAdmin}, nil
}

func (s *Service) IsTokenInvalid(token string) (bool, error) {
	return s.repo.IsTokenInvalid(token)
}
