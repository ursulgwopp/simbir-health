package repository

import (
	"slices"

	"github.com/lib/pq"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/models"
)

func (r *PostgresRepository) SignUp(req models.SignUpRequest) (int, error) {
	// CHECKING IF USERNAME EXISTS
	exists, err := CheckUsernameExists(r.db, req.Username)
	if err != nil {
		return -1, err
	}

	if exists {
		return -1, custom_errors.ErrUsernameExists
	}

	// INSERTING NEW ACCOUNT INTO TABLE
	var id int
	query := `INSERT INTO accounts (last_name, first_name, username, hash_password) VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.db.QueryRow(query, req.LastName, req.FirstName, req.Username, req.Password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *PostgresRepository) SignIn(req models.SignInRequest) (models.TokenInfo, error) {
	// SIGNING IN TO ACCOUNT
	var id int
	var roles []string
	query := `SELECT id, roles FROM accounts WHERE username = $1 AND hash_password = $2 AND is_deleted = false`

	row := r.db.QueryRow(query, req.Username, req.Password)
	if err := row.Scan(&id, pq.Array(&roles)); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.TokenInfo{}, custom_errors.ErrSignIn
		}

		return models.TokenInfo{}, err
	}

	// SETTING TOKEN INFO
	isAdmin := slices.Contains(roles, "Admin")
	isManager := slices.Contains(roles, "Manager")
	isDoctor := slices.Contains(roles, "Doctor")

	return models.TokenInfo{UserId: id, IsAdmin: isAdmin, IsManager: isManager, IsDoctor: isDoctor}, nil
}

func (r *PostgresRepository) SignOut(token string) error {
	query := `INSERT INTO blacklist (token) VALUES ($1)`
	_, err := r.db.Exec(query, token)

	return err
}

func (r *PostgresRepository) Refresh(token string) error {
	query := `INSERT INTO blacklist (token) VALUES ($1)`
	_, err := r.db.Exec(query, token)

	return err
}

func (r *PostgresRepository) IsTokenInvalid(token string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM blacklist WHERE token = $1)"

	err := r.db.QueryRow(query, token).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
