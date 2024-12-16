package service

import "github.com/ursulgwopp/simbir-health/internal/account_microservice/models"

func (s *Service) UserGetAccount(accountId int) (models.AccountResponse, error) {
	return s.repo.UserGetAccount(accountId)
}

func (s *Service) UserUpdateAccount(accountId int, req models.AccountUpdate) error {
	req.Password = generatePasswordHash(req.Password)
	return s.repo.UserUpdateAccount(accountId, req)
}

func (s *Service) UserListDoctors(nameFilter string, from int, count int) ([]models.DoctorResponse, error) {
	return s.repo.UserListDoctors(nameFilter, from, count)
}

func (s *Service) UserGetDoctor(doctorId int) (models.DoctorResponse, error) {
	return s.repo.UserGetDoctor(doctorId)
}
