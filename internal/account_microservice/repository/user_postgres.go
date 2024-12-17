package repository

import (
	"github.com/lib/pq"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
)

func (r *PostgresRepository) UserGetAccount(accountId int) (models.AccountResponse, error) {
	// CHECKING IF ID EXISTS
	exists, err := CheckIdExists(r.db, accountId)
	if err != nil {
		return models.AccountResponse{}, err
	}

	if !exists {
		return models.AccountResponse{}, custom_errors.ErrIdNotFound

	}

	// GETTING DATA FROM DB
	var account models.AccountResponse
	query := `SELECT last_name, first_name, username, roles FROM accounts WHERE id = $1 AND is_deleted = false`

	row := r.db.QueryRow(query, accountId)
	if err := row.Scan(&account.LastName, &account.FirstName, &account.Username, pq.Array(&account.Roles)); err != nil {
		return models.AccountResponse{}, err
	}

	account.Id = accountId
	return account, nil
}

func (r *PostgresRepository) UserUpdateAccount(accountId int, req models.AccountUpdate) error {
	// CHECKING IF ID EXISTS
	exists, err := CheckIdExists(r.db, accountId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound

	}

	// UPDATING ACCOUNT
	query := `UPDATE accounts SET last_name = $1, first_name = $2, hash_password = $3 WHERE id = $4 AND is_deleted = false`
	_, err = r.db.Exec(query, req.LastName, req.FirstName, req.Password, accountId)

	return err
}

func (r *PostgresRepository) UserListDoctors(nameFilter string, from int, count int) ([]models.DoctorResponse, error) {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO FILTER BY NAME
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// GETTING DOCTORS FROM DB
	query := `SELECT id, last_name, first_name, username FROM accounts WHERE 'Doctor' = ANY(roles) AND is_deleted = false ORDER BY id OFFSET $1 LIMIT $2`

	rows, err := r.db.Query(query, from, count)
	if err != nil {
		return []models.DoctorResponse{}, err
	}
	defer rows.Close()

	var doctors []models.DoctorResponse
	for rows.Next() {
		doctor := models.DoctorResponse{}
		if err := rows.Scan(&doctor.Id, &doctor.LastName, &doctor.FirstName, &doctor.Username); err != nil {
			return nil, err
		}

		doctors = append(doctors, doctor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return doctors, nil
}

func (r *PostgresRepository) UserGetDoctor(doctorId int) (models.DoctorResponse, error) {
	// CHECKING IF ID EXISTS
	exists, err := CheckDoctorIdExists(r.db, doctorId)
	if err != nil {
		return models.DoctorResponse{}, err
	}

	if !exists {
		return models.DoctorResponse{}, custom_errors.ErrIdNotFound

	}

	// GETTNG DOCTOR FROM DB
	var doctor models.DoctorResponse
	query := `SELECT last_name, first_name, username FROM accounts WHERE 'Doctor' = ANY(roles) and id = $1 AND is_deleted = false`

	row := r.db.QueryRow(query, doctorId)
	if err := row.Scan(&doctor.LastName, &doctor.FirstName, &doctor.Username); err != nil {
		return models.DoctorResponse{}, err
	}

	doctor.Id = doctorId
	return doctor, nil
}
