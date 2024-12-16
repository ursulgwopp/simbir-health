package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

// @Summary GetAccount
// @Tags User
// @Security ApiKeyAuth
// @Description Get My Account
// @ID get-my-account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.AccountResponse
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Accounts/Me [get]
func (t *Transport) userGetAccount(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	account, err := t.service.UserGetAccount(id)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Summary UpdateAccount
// @Tags User
// @Security ApiKeyAuth
// @Description Update My Account
// @ID update-my-account
// @Accept  json
// @Produce  json
// @Param Input body models.AccountUpdate true "Account Update"
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Accounts/Update [put]
func (t *Transport) userUpdateAccount(c *gin.Context) {
	var req models.AccountUpdate

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.UserUpdateAccount(id, req); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "ok"})
}

// @Summary List Doctors
// @Tags User
// @Security ApiKeyAuth
// @Description List Doctors
// @ID list-doctors
// @Accept  json
// @Produce  json
// @Param from query int true "from"
// @Param count query int true "count"
// @Success 200 {array} models.DoctorResponse
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Doctors [get]
func (t *Transport) userListDoctors(c *gin.Context) {
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

	doctors, err := t.service.UserListDoctors("", from, count)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, doctors)
}

// @Summary GetDoctor
// @Tags User
// @Security ApiKeyAuth
// @Description Get Doctor
// @ID get-doctor
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.DoctorResponse
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Doctors/{id} [get]
func (t *Transport) userGetDoctor(c *gin.Context) {
	doctorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	doctor, err := t.service.UserGetDoctor(doctorId)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, doctor)
}
