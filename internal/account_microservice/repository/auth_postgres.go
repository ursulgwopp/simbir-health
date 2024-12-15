package repository

import (
	"slices"

	"github.com/lib/pq"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

func (r *PostgresRepository) SignUp(req models.SignUpRequest) (int, error) {
	var id int
	query := `INSERT INTO accounts (last_name, first_name, username, hash_password) VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.db.QueryRow(query, req.LastName, req.FirstName, req.Username, req.Password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) SignIn(req models.SignInRequest) (models.TokenInfo, error) {
	var id int
	var roles []string
	query := `SELECT id, roles FROM accounts WHERE username = $1 AND hash_password = $2`

	row := r.db.QueryRow(query, req.Username, req.Password)
	if err := row.Scan(&id, pq.Array(&roles)); err != nil {
		return models.TokenInfo{}, err
	}

	isAdmin := slices.Contains(roles, "admin")

	return models.TokenInfo{UserId: id, IsAdmin: isAdmin}, nil
}

func (r *PostgresRepository) SignOut(token string) error {
	query := `INSERT INTO blacklist (token) VALUES ($1)`
	_, err := r.db.Exec(query, token)

	return err
}
