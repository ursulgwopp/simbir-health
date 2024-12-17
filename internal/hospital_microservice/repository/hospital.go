package repository

import "github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"

// CreateHospital implements service.HospitalRepository.
func (*PostgresRepository) CreateHospital(req models.HospitalRequest) (int, error) {
	panic("unimplemented")
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
