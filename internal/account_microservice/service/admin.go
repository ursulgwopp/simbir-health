package service

import (
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

func (s *Service) AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error) {
	// VALIDATING QUERY PARAMS
	if from < 0 {
		return []models.AdminAccountResponse{}, custom_errors.ErrInvalidFrom
	}

	if count < 0 {
		return []models.AdminAccountResponse{}, custom_errors.ErrInvalidCount
	}

	// PASSING PARAMS TO REPOSITORY LAYER
	return s.repo.AdminListAccounts(from, count)
}

func (s *Service) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
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
	return s.repo.AdminCreateAccount(req)
}

func (s *Service) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	// VALIDATING REQUEST DATA
	if err := validateName(req.LastName, req.FirstName); err != nil {
		return err
	}

	if err := validateUsername(req.Username); err != nil {
		return err
	}

	if err := validatePassword(req.Password); err != nil {
		return err
	}

	// HASHING PASSWORD\
	req.Password = generatePasswordHash(req.Password)

	// PASSING VALIDATED DATA TO REPOSITORY LAYER
	return s.repo.AdminUpdateAccount(accountId, req)
}

func (s *Service) AdminDeleteAccount(accountId int) error {
	return s.repo.AdminDeleteAccount(accountId)
}
