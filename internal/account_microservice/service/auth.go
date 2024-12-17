package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ursulgwopp/simbir-health/internal/models"
)

func (s *Service) SignUp(req models.SignUpRequest) (int, error) {
	// VALIDATING REQUEST DATA
	if err := validateName(req.LastName, req.FirstName); err != nil {
		return -1, err
	}

	if err := validateUsername(req.Username); err != nil {
		return -1, err
	}

	if err := validatePassword(req.Password); err != nil {
		return -1, err
	}

	// HASHING PASSWORD
	req.Password = generatePasswordHash(req.Password)

	// PASSING VALIDATED DATA TO REPOSITORY LAYER
	return s.repo.SignUp(req)
}

func (s *Service) SignIn(req models.SignInRequest) (string, error) {
	// VALIDATING REQUEST DATA
	if err := validateUsername(req.Username); err != nil {
		return "", err
	}

	if err := validatePassword(req.Password); err != nil {
		return "", err
	}

	// HASHING PASSWORD
	req.Password = generatePasswordHash(req.Password)

	// PASSING VALIDATED DATA TO REPOSITORY LAYER
	tokenInfo, err := s.repo.SignIn(req)
	if err != nil {
		return "", err
	}

	// CREATING JWT TOKEN
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
	// PASSING VALIDATED DATA TO REPOSITORY LAYER
	return s.repo.SignOut(token)
}
