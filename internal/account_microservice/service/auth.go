package service

import (
	"errors"
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
