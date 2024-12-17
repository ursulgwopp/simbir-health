package service

import "github.com/ursulgwopp/simbir-health/internal/models"

type AccountRepository interface {
	SignUp(req models.SignUpRequest) (int, error)
	SignIn(req models.SignInRequest) (models.TokenInfo, error)
	SignOut(token string) error

	Refresh(token string) error
	IsTokenInvalid(token string) (bool, error)

	UserGetAccount(accountId int) (models.AccountResponse, error)
	UserUpdateAccount(accountId int, req models.AccountUpdate) error
	UserListDoctors(nameFilter string, from int, count int) ([]models.DoctorResponse, error)
	UserGetDoctor(doctorId int) (models.DoctorResponse, error)

	AdminListAccounts(from int, count int) ([]models.AdminAccountResponse, error)
	AdminCreateAccount(req models.AdminAccountRequest) (int, error)
	AdminUpdateAccount(accountId int, req models.AdminAccountRequest) error
	AdminDeleteAccount(accountId int) error
}

type Service struct {
	repo AccountRepository
}

func NewService(repo AccountRepository) *Service {
	return &Service{
		repo: repo,
	}
}
