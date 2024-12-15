package repository

import "github.com/ursulgwopp/simbir-health/internal/account_microservice/models"

func (r *PostgresRepository) AdminListAccounts(from int, count int) ([]models.AccountResponse, error) {
	return []models.AccountResponse{}, nil
}

func (r *PostgresRepository) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
	return -1, nil
}

func (r *PostgresRepository) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	return nil
}

func (r *PostgresRepository) AdminDeleteAccount(accountId int) error {
	return nil
}
