package service

import "github.com/ursulgwopp/simbir-health/internal/account_microservice/models"

func (s *Service) AdminListAccounts(from int, count int) ([]models.AccountResponse, error) {
	return s.repo.AdminListAccounts(from, count)
}

func (s *Service) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
	return s.repo.AdminCreateAccount(req)
}

func (s *Service) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	return s.repo.AdminUpdateAccount(accountId, req)
}

func (s *Service) AdminDeleteAccount(accountId int) error {
	return s.repo.AdminDeleteAccount(accountId)
}
