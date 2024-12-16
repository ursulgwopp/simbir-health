package repository

import "github.com/jmoiron/sqlx"

func CheckUsernameExists(db *sqlx.DB, username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM accounts WHERE username = $1)"
	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
