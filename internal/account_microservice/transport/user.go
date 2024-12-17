package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/models"
)

// @Router /api/Accounts/Me [get]
// @Summary GetAccount
// @Tags User
// @Security ApiKeyAuth
// @Description Get My Account
// @ID get-my-account
// @Accept json
// @Produce json
// @Success 200 {object} models.AccountResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) userGetAccount(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	account, err := t.service.UserGetAccount(id)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserIdNotFound) {
			models.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, account)
}

// @Router /api/Accounts/Update [put]
// @Summary UpdateAccount
// @Tags User
// @Security ApiKeyAuth
// @Description Update My Account
// @ID update-my-account
// @Accept json
// @Produce json
// @Param Input body models.AccountUpdate true "Account Update"
// @Success 200 {object} models.Response
// @Failure 400,404,409 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) userUpdateAccount(c *gin.Context) {
	// UNMARSHALLING REQUEST BODY
	var req models.AccountUpdate

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// GETTING USER ID
	id, err := getUserId(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// PASSING DATA TO SERVICE LAYER
	if err := t.service.UserUpdateAccount(id, req); err != nil {
		if errors.Is(err, custom_errors.ErrFirstNameInvalid) ||
			errors.Is(err, custom_errors.ErrLastNameInvalid) ||
			errors.Is(err, custom_errors.ErrShortPassword) ||
			errors.Is(err, custom_errors.ErrPasswordWithoutDigits) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, custom_errors.ErrUserIdNotFound) {
			models.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "account successfully updated"})
}

// @Router /api/Doctors [get]
// @Summary List Doctors
// @Tags User
// @Security ApiKeyAuth
// @Description List Doctors
// @ID list-doctors
// @Accept json
// @Produce json
// @Param from query int true "From"
// @Param count query int true "Count"
// @Success 200 {array} models.DoctorResponse
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) userListDoctors(c *gin.Context) {
	// GETTING QUERY PARAMS
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

	// PASSING PARAMS TO SERVICE LAYER
	doctors, err := t.service.UserListDoctors("", from, count)
	if err != nil {
		if errors.Is(err, custom_errors.ErrInvalidFrom) ||
			errors.Is(err, custom_errors.ErrInvalidCount) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, doctors)
}

// @Router /api/Doctors/{id} [get]
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
func (t *Transport) userGetDoctor(c *gin.Context) {
	// GETTING DATAA
	doctorId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// PASSING DATA TO SERVICE
	doctor, err := t.service.UserGetDoctor(doctorId)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserIdNotFound) {
			models.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, doctor)
}
