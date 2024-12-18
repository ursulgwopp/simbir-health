package service

import (
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"
)

func (s *Service) CreateHospital(req models.HospitalRequest) (int, error) {
	if err := validateName(req.Name); err != nil {
		return -1, custom_errors.ErrInvalidName
	}

	if err := validateAddress(req.Address); err != nil {
		return -1, custom_errors.ErrInvalidAddress
	}

	if err := validatePhone(req.ContactPhone); err != nil {
		return -1, custom_errors.ErrInvalidPhone
	}

	return s.repo.CreateHospital(req)
}

func (s *Service) DeleteHospital(hospitalId int) error {
	return s.repo.DeleteHospital(hospitalId)
}

func (s *Service) GetHospital(hospitalId int) (models.HospitalResponse, error) {
	return s.repo.GetHospital(hospitalId)
}

func (s *Service) GetHospitalRooms(hospitalId int) ([]string, error) {
	return s.repo.GetHospitalRooms(hospitalId)
}

func (s *Service) ListHospitals(from int, count int) ([]models.HospitalResponse, error) {
	return s.repo.ListHospitals(from, count)
}

// UpdateHospital implements transport.HospitalService.
func (s *Service) UpdateHospital(hospitalId int, req models.HospitalRequest) error {
	if err := validateName(req.Name); err != nil {
		return custom_errors.ErrInvalidName
	}

	if err := validateAddress(req.Address); err != nil {
		return custom_errors.ErrInvalidAddress
	}

	if err := validatePhone(req.ContactPhone); err != nil {
		return custom_errors.ErrInvalidPhone
	}

	return s.repo.UpdateHospital(hospitalId, req)
}
