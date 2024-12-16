package repository

import (
	"github.com/lib/pq"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

func (r *PostgresRepository) AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error) {
	query := `SELECT id, last_name, first_name, username, roles, is_deleted FROM accounts ORDER BY id OFFSET $1 LIMIT $2`

	rows, err := r.db.Query(query, from, count)
	if err != nil {
		return []models.AdminAccountResponse{}, err
	}
	defer rows.Close()

	var accounts []models.AdminAccountResponse
	for rows.Next() {
		account := models.AdminAccountResponse{}
		if err := rows.Scan(&account.Id, &account.LastName, &account.FirstName, &account.Username, pq.Array(&account.Roles), &account.IsDeleted); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *PostgresRepository) AdminCreateAccount(req models.AdminAccountRequest) (int, error) {
	var id int
	query := `INSERT INTO accounts (last_name, first_name, username, hash_password, roles) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	row := r.db.QueryRow(query, req.LastName, req.FirstName, req.Username, req.Password, pq.Array(req.Roles))
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error {
	query := `UPDATE accounts SET last_name = $1, first_name = $2, username = $3, hash_password = $4, roles = $5 WHERE id = $6`

	_, err := r.db.Exec(query, req.LastName, req.FirstName, req.Username, req.Password, pq.Array(req.Roles), accountId)
	return err
}

func (r *PostgresRepository) AdminDeleteAccount(accountId int) error {
	query := `UPDATE accounts SET is_deleted = true WHERE id = $1`

	_, err := r.db.Exec(query, accountId)
	return err
}
