package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"
)

// @Router /api/Hospital [post]
// @Summary AddHospital
// @Security ApiKeyAuth
// @Tags Hospital
// @Description Add Hospital
// @ID add-hospital
// @Accept json
// @Produce json
// @Param Input body models.HospitalRequest true "Hospital Info"
// @Success 200 {array} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) createHospital(c *gin.Context) {
	var req models.HospitalRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.CreateHospital(req)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]int{"id": id})
}
