package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

// @Summary AdminListAccounts
// @Security ApiKeyAuth
// @Tags Admin
// @Description Admin List Accounts
// @ID admin-list-accounts
// @Accept  json
// @Produce  json
// @Param from query int true "from"
// @Param count query int true "count"
// @Success 200 {array} models.AdminAccountRequest
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Accounts [get]
func (t *Transport) adminListAccounts(c *gin.Context) {
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

	accounts, err := t.service.AdminListAccounts(from, count)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, accounts)
}

// @Summary AdminCreateAccount
// @Security ApiKeyAuth
// @Tags Admin
// @Description Admin Create Account
// @ID admin-create-account
// @Accept  json
// @Produce  json
// @Param Input body models.AdminAccountRequest true "Account Info"
// @Success 201 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Accounts [post]
func (t *Transport) adminCreateAccount(c *gin.Context) {
	var req models.AdminAccountRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.AdminCreateAccount(req)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.Response{Message: id})
}

// @Summary AdminUpdateAccount
// @Security ApiKeyAuth
// @Tags Admin
// @Description Admin Update Account
// @ID admin-update-account
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Param Input body models.AdminAccountRequest true "Account Info"
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Accounts/{id} [put]
func (t *Transport) adminUpdateAccount(c *gin.Context) {
	var req models.AdminAccountRequest

	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.AdminUpdateAccount(accountId, req); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.Response{Message: "ok"})
}

// @Summary AdminDeleteAccount
// @Security ApiKeyAuth
// @Tags Admin
// @Description Admin Delete Account
// @ID admin-delete-account
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Accounts/{id} [delete]
func (t *Transport) adminDeleteAccount(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.AdminDeleteAccount(accountId); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "ok"})
}
