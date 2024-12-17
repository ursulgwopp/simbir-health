package repository

import "github.com/jmoiron/sqlx"

func CheckIdExists(db *sqlx.DB, id int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM hospitals WHERE id = $1 AND is_deleted = false)"
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
