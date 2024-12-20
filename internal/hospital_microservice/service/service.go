package service

import (
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"
)

type HospitalRepository interface {
	ListHospitals(from int, count int) ([]models.HospitalResponse, error)
	GetHospital(hospitalId int) (models.HospitalResponse, error)
	GetHospitalRooms(hospitalId int) ([]string, error)
	CreateHospital(req models.HospitalRequest) (int, error)
	UpdateHospital(hospitalId int, req models.HospitalRequest) error
	DeleteHospital(hospitalId int) error
}

type Service struct {
	repo HospitalRepository
}

func NewService(repo HospitalRepository) *Service {
	return &Service{
		repo: repo,
	}
}
