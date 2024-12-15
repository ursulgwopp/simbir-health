package repository

// func (r *PostgresRepository) Validate(token string) (models.AccountResponse, error) {
// 	return models.AccountResponse{}, nil
// }

func (r *PostgresRepository) Refresh(token string) error {
	return nil
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
