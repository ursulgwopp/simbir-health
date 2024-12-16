package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

func (s *Service) SignUp(req models.SignUpRequest) (int, error) {
	req.Password = generatePasswordHash(req.Password)
	return s.repo.SignUp(req)
}

func (s *Service) SignIn(req models.SignInRequest) (string, error) {
	req.Password = generatePasswordHash(req.Password)

	tokenInfo, err := s.repo.SignIn(req)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:    tokenInfo.UserId,
		IsAdmin:   tokenInfo.IsAdmin,
		IsManager: tokenInfo.IsManager,
		IsDoctor:  tokenInfo.IsDoctor,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *Service) SignOut(token string) error {
	return s.repo.SignOut(token)
}
