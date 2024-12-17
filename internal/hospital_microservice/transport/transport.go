package transport

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ursulgwopp/simbir-health/docs"
	"github.com/ursulgwopp/simbir-health/internal/models"
)

type HospitalService interface {
	ListHospitals(from int, count int) ([]models.HospitalResponse, error)
	GetHospital(hospitalId int) (models.HospitalResponse, error)
	GetHospitalRooms(hospitalId int) ([]string, error)
	CreateHospital(req models.HospitalRequest) (int, error)
	UpdateHospital(hospitalId int, req models.HospitalResponse) error
	DeleteHospital(hospitalId int) error
}

type Transport struct {
	service HospitalService
}

func NewTransport(service HospitalService) *Transport {
	return &Transport{service: service}
}

func (t *Transport) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		hospitals := api.Group("/Hospitals")
		{
			hospitals.GET("/", nil)
			hospitals.GET("/:id", nil)
			hospitals.GET("/:id/Rooms", nil)
			hospitals.POST("/", nil)
			hospitals.PUT("/:id", nil)
			hospitals.DELETE("/:id", nil)
		}
	}

	return router
}
