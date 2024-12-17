package service

import "github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"

// CreateHospital implements transport.HospitalService.
func (*Service) CreateHospital(req models.HospitalRequest) (int, error) {
	panic("unimplemented")
}

// DeleteHospital implements transport.HospitalService.
func (*Service) DeleteHospital(hospitalId int) error {
	panic("unimplemented")
}

// GetHospital implements transport.HospitalService.
func (*Service) GetHospital(hospitalId int) (models.HospitalResponse, error) {
	panic("unimplemented")
}

// GetHospitalRooms implements transport.HospitalService.
func (*Service) GetHospitalRooms(hospitalId int) ([]string, error) {
	panic("unimplemented")
}

// ListHospitals implements transport.HospitalService.
func (*Service) ListHospitals(from int, count int) ([]models.HospitalResponse, error) {
	panic("unimplemented")
}

// UpdateHospital implements transport.HospitalService.
func (*Service) UpdateHospital(hospitalId int, req models.HospitalResponse) error {
	panic("unimplemented")
}
