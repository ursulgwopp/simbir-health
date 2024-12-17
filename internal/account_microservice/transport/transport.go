package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/models"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ursulgwopp/simbir-health/docs"
)

type AccountService interface {
	SignUp(req models.SignUpRequest) (int, error)
	SignIn(req models.SignInRequest) (string, error)
	SignOut(token string) error

	Validate(token string) (models.TokenInfo, error)
	Refresh(token string) (string, error)
	ParseToken(token string) (models.TokenInfo, error)
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

type Transport struct {
	service AccountService
}

func NewTransport(service AccountService) *Transport {
	return &Transport{service: service}
}

func (t *Transport) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		Authentication := api.Group("/Authentication")
		{
			Authentication.POST("/SignUp", t.signUp)
			Authentication.POST("/SignIn", t.signIn)
			Authentication.PUT("/SignOut", t.userIdentity, t.signOut)

			Authentication.GET("/Validate", t.validate)
			Authentication.POST("/Refresh", t.refresh)
		}

		Accounts := api.Group("/Accounts", t.userIdentity)
		{
			Accounts.GET("/Me", t.userGetAccount)
			Accounts.PUT("/Update", t.userUpdateAccount)

			Accounts.GET("/", t.adminIdentity, t.adminListAccounts)
			Accounts.POST("/", t.adminIdentity, t.adminCreateAccount)
			Accounts.PUT("/:id", t.adminIdentity, t.adminUpdateAccount)
			Accounts.DELETE("/:id", t.adminIdentity, t.adminDeleteAccount)
		}

		Doctors := api.Group("/Doctors", t.userIdentity)
		{
			Doctors.GET("/", t.userListDoctors)
			Doctors.GET("/:id", t.userGetDoctor)
		}
	}

	return router
}
