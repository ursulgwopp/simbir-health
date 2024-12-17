package repository

import (
	"github.com/lib/pq"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"
)

// CreateHospital implements service.HospitalRepository.
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
func (*PostgresRepository) DeleteHospital(hospitalId int) error {
	panic("unimplemented")
}

// GetHospital implements service.HospitalRepository.
func (*PostgresRepository) GetHospital(hospitalId int) (models.HospitalResponse, error) {
	panic("unimplemented")
}

// GetHospitalRooms implements service.HospitalRepository.
func (*PostgresRepository) GetHospitalRooms(hospitalId int) ([]string, error) {
	panic("unimplemented")
}

// ListHospitals implements service.HospitalRepository.
func (*PostgresRepository) ListHospitals(from int, count int) ([]models.HospitalResponse, error) {
	panic("unimplemented")
}

// UpdateHospital implements service.HospitalRepository.
func (*PostgresRepository) UpdateHospital(hospitalId int, req models.HospitalResponse) error {
	panic("unimplemented")
}
