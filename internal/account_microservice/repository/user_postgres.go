package repository

import "github.com/ursulgwopp/simbir-health/internal/account_microservice/models"

func (r *PostgresRepository) UserGetAccount(accountId int) (models.AccountResponse, error) {
	return models.AccountResponse{}, nil
}

func (r *PostgresRepository) UserUpdateAccount(accountId int, req models.AccountUpdate) error {
	return nil
}

func (r *PostgresRepository) UserListDoctors(nameFilter string, from int, count int) ([]models.DoctorResponse, error) {
	return []models.DoctorResponse{}, nil
}

func (r *PostgresRepository) UserGetDoctor(doctorId int) (models.DoctorResponse, error) {
	return models.DoctorResponse{}, nil
}
