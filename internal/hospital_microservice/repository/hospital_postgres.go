package repository

import (
	"github.com/lib/pq"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"
)

func (r *PostgresRepository) CreateHospital(req models.HospitalRequest) (int, error) {
	// INSERTING ACCOUNT INTO ACCOUNTS
	var id int
	query := `INSERT INTO hospitals (name, address, contact_phone, rooms) VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.db.QueryRow(query, req.Name, req.Address, req.ContactPhone, pq.Array(req.Rooms))
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

// DeleteHospital implements service.HospitalRepository.
func (r *PostgresRepository) DeleteHospital(hospitalId int) error {
	// CHECKING IF ID EXISTS
	exists, err := CheckIdExists(r.db, hospitalId)
	if err != nil {
		return err
	}

	if !exists {
		return custom_errors.ErrIdNotFound

	}

	// SOFT DELETING HOPSITAL
	query := `UPDATE hospitals SET is_deleted = true WHERE id = $1`

	_, err = r.db.Exec(query, hospitalId)
	return err
}

// GetHospital implements service.HospitalRepository.
func (r *PostgresRepository) GetHospital(hospitalId int) (models.HospitalResponse, error) {
	exists, err := CheckIdExists(r.db, hospitalId)
	if err != nil {
		return models.HospitalResponse{}, err
	}

	if !exists {
		return models.HospitalResponse{}, custom_errors.ErrIdNotFound
	}

	var hospital models.HospitalResponse
	query := `SELECT id, name, address, contact_phone FROM hospitals WHERE id = $1`

	row := r.db.QueryRow(query, hospitalId)
	if err := row.Scan(&hospital.Id, &hospital.Name, &hospital.Address, &hospital.ContactPhone); err != nil {
		return models.HospitalResponse{}, err
	}

	return hospital, nil
}

func (r *PostgresRepository) GetHospitalRooms(hospitalId int) ([]string, error) {
	exists, err := CheckIdExists(r.db, hospitalId)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, custom_errors.ErrIdNotFound
	}

	var rooms []string
	query := `SELECT rooms FROM hospitals WHERE id = $1`

	row := r.db.QueryRow(query, hospitalId)
	if err := row.Scan(pq.Array(&rooms)); err != nil {
		return nil, err
	}

	return rooms, nil
}

// ListHospitals implements service.HospitalRepository.
func (r *PostgresRepository) ListHospitals(from int, count int) ([]models.HospitalResponse, error) {
	query := `SELECT id, name, address, contact_phone FROM hospitals WHERE is_deleted = false ORDER BY id OFFSET $1 LIMIT $2`

	rows, err := r.db.Query(query, from, count)
	if err != nil {
		return []models.HospitalResponse{}, err
	}
	defer rows.Close()

	var hospitals []models.HospitalResponse
	for rows.Next() {
		var hospital models.HospitalResponse

		if err := rows.Scan(&hospital.Id, &hospital.Name, &hospital.Address, &hospital.ContactPhone); err != nil {
			return []models.HospitalResponse{}, err
		}

		hospitals = append(hospitals, hospital)
	}

	return hospitals, nil
}

// UpdateHospital implements service.HospitalRepository.
func (*PostgresRepository) UpdateHospital(hospitalId int, req models.HospitalResponse) error {
	panic("unimplemented")
}
