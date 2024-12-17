package transport

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
)

// @Router /api/Accounts [get]
// @Summary AdminListAccounts
// @Security ApiKeyAuth
// @Tags Admin
// @Description Admin List Accounts
// @ID admin-list-accounts
// @Accept json
// @Produce json
// @Param from query int true "From"
// @Param count query int true "Count"
// @Success 200 {array} models.AdminAccountRequest
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminListAccounts(c *gin.Context) {
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
	accounts, err := t.service.AdminListAccounts(from, count)
	if err != nil {
		if errors.Is(err, custom_errors.ErrInvalidFrom) ||
			errors.Is(err, custom_errors.ErrInvalidCount) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// @Router /api/Accounts [post]
// @Summary AdminCreateAccount
// @Security ApiKeyAuth
// @Tags Admin
// @Description Admin Create Account
// @ID admin-create-account
// @Accept json
// @Produce json
// @Param Input body models.AdminAccountRequest true "Account Info"
// @Success 201 {object} models.Response
// @Failure 400,409 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminCreateAccount(c *gin.Context) {
	// UNMARSHALLING REQUEST BODY
	var req models.AdminAccountRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// PASSING BODY TO SERVICE LAYER
	id, err := t.service.AdminCreateAccount(req)
	if err != nil {
		if errors.Is(err, custom_errors.ErrFirstNameInvalid) ||
			errors.Is(err, custom_errors.ErrLastNameInvalid) ||
			errors.Is(err, custom_errors.ErrUsernameInvalidCharacters) ||
			errors.Is(err, custom_errors.ErrUsernameInvalidLength) ||
			errors.Is(err, custom_errors.ErrShortPassword) ||
			errors.Is(err, custom_errors.ErrPasswordWithoutDigits) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, custom_errors.ErrUsernameExists) {
			models.NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// @Router /api/Accounts/{id} [put]
// @Summary AdminUpdateAccount
// @Security ApiKeyAuth
// @Tags Admin
// @Description Admin Update Account
// @ID admin-update-account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param Input body models.AdminAccountRequest true "Account Info"
// @Success 200 {object} models.Response
// @Failure 400,404,409 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminUpdateAccount(c *gin.Context) {
	// GETTING PATH PARAMS
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// UNMARSHALLING REQUEST BODY
	var req models.AdminAccountRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// PASSSING DATA TO SERVICE LAYER
	if err := t.service.AdminUpdateAccount(accountId, req); err != nil {
		if errors.Is(err, custom_errors.ErrFirstNameInvalid) ||
			errors.Is(err, custom_errors.ErrLastNameInvalid) ||
			errors.Is(err, custom_errors.ErrUsernameInvalidCharacters) ||
			errors.Is(err, custom_errors.ErrUsernameInvalidLength) ||
			errors.Is(err, custom_errors.ErrShortPassword) ||
			errors.Is(err, custom_errors.ErrPasswordWithoutDigits) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, custom_errors.ErrUsernameExists) {
			models.NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.Response{Message: "account updated successfully"})
}

// @Router /api/Accounts/{id} [delete]
// @Summary AdminDeleteAccount
// @Security ApiKeyAuth
// @Tags Admin
// @Description Admin Delete Account
// @ID admin-delete-account
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) adminDeleteAccount(c *gin.Context) {
	// GETTING PATH PARAM
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// PASSING PARAM TO SERVICE LAYER
	if err := t.service.AdminDeleteAccount(accountId); err != nil {
		if errors.Is(err, custom_errors.ErrIdNotFound) {
			models.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "account successfully deleted"})
}
