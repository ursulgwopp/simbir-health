package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"
)

// @Router /api/Hospital [get]
// @Summary ListHospitals
// @Security ApiKeyAuth
// @Tags Hospital
// @Description List Hospitals
// @ID list-hospitals
// @Accept json
// @Produce json
// @Param from query int true "From"
// @Param count query int true "Count"
// @Success 200 {array} models.HospitalResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) listHospitals(c *gin.Context) {
	from_ := c.Query("from")
	count_ := c.Query("count")

	from, err := strconv.Atoi(from_)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	count, err := strconv.Atoi(count_)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hospitals, err := t.service.ListHospitals(from, count)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, hospitals)
}

// @Router /api/Hospital/{id} [get]
// @Summary GetHospital
// @Security ApiKeyAuth
// @Tags Hospital
// @Description Get Hospital
// @ID get-hospital
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.HospitalResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) getHospital(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hospital, err := t.service.GetHospital(id)
	if err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, hospital)
}

// @Router /api/Hospital/{id}/Rooms [get]
// @Summary GetHospitalRooms
// @Security ApiKeyAuth
// @Tags Hospital
// @Description Get Hospital Rooms
// @ID get-hospital-rooms
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {array} string
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) getHospitalRooms(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rooms, err := t.service.GetHospitalRooms(id)
	if err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rooms)
}

// @Router /api/Hospital [post]
// @Summary AddHospital
// @Security ApiKeyAuth
// @Tags Hospital
// @Description Add Hospital
// @ID add-hospital
// @Accept json
// @Produce json
// @Param Input body models.HospitalRequest true "Hospital Info"
// @Success 200 {object} models.Response
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
		if errors.Is(err, custom_errors.ErrInvalidName) ||
			errors.Is(err, custom_errors.ErrInvalidAddress) ||
			errors.Is(err, custom_errors.ErrInvalidPhone) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]int{"id": id})
}

// @Router /api/Hospital/{id} [delete]
// @Summary DeleteHospital
// @Security ApiKeyAuth
// @Tags Hospital
// @Description Delete Hospital
// @ID delete-hospital
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) deleteHospital(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.DeleteHospital(id); err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "hospital successfully deleted")
}

// @Router /api/Hospital/{id} [put]
// @Summary UpdateHospital
// @Security ApiKeyAuth
// @Tags Hospital
// @Description Update Hospital
// @ID update-hospital
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param Input body models.HospitalRequest true "Hospital Info"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) updateHospital(c *gin.Context) {
	id, err := parseId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var req models.HospitalRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.UpdateHospital(id, req); err != nil {
		if errors.Is(err, custom_errors.ErrInvalidName) ||
			errors.Is(err, custom_errors.ErrInvalidAddress) ||
			errors.Is(err, custom_errors.ErrInvalidPhone) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "hospital successfully updated")
}
