package repository

import (
	"github.com/lib/pq"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

func (r *PostgresRepository) UserGetAccount(accountId int) (models.AccountResponse, error) {
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
	query := `UPDATE accounts SET last_name = $1, first_name = $2, hash_password = $3 WHERE id = $4 AND is_deleted = false`
	_, err := r.db.Exec(query, req.LastName, req.FirstName, req.Password, accountId)

	return err
}

func (r *PostgresRepository) UserListDoctors(nameFilter string, from int, count int) ([]models.DoctorResponse, error) {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// TODO FILTER BY NAME
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
	var doctor models.DoctorResponse
	query := `SELECT last_name, first_name, username FROM accounts WHERE 'Doctor' = ANY(roles) and id = $1 AND is_deleted = false`

	row := r.db.QueryRow(query, doctorId)
	if err := row.Scan(&doctor.LastName, &doctor.FirstName, &doctor.Username); err != nil {
		return models.DoctorResponse{}, err
	}

	doctor.Id = doctorId
	return doctor, nil
}
