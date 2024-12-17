package service

import (
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
)

func (s *Service) UserGetAccount(accountId int) (models.AccountResponse, error) {
	return s.repo.UserGetAccount(accountId)
}

func (s *Service) UserUpdateAccount(accountId int, req models.AccountUpdate) error {
	// VALIDATING REQUEST DATA
	if err := validateName(req.LastName, req.FirstName); err != nil {
		return err
	}

	if err := validatePassword(req.Password); err != nil {
		return err
	}

	// HASHING PASSWORD
	req.Password = generatePasswordHash(req.Password)

	// PASSING VALIDATED DATA TO REPOSITORY LAYER
	req.Password = generatePasswordHash(req.Password)
	return s.repo.UserUpdateAccount(accountId, req)
}

func (s *Service) UserListDoctors(nameFilter string, from int, count int) ([]models.DoctorResponse, error) {
	// VALIDATING QUERY PARAMS
	if from < 0 {
		return []models.DoctorResponse{}, custom_errors.ErrInvalidFrom
	}

	if count < 0 {
		return []models.DoctorResponse{}, custom_errors.ErrInvalidCount
	}

	return s.repo.UserListDoctors(nameFilter, from, count)
}

func (s *Service) UserGetDoctor(doctorId int) (models.DoctorResponse, error) {
	return s.repo.UserGetDoctor(doctorId)
}
