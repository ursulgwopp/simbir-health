package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

// @Summary Validate
// @Tags Authentication
// @Description Validate token
// @ID validate
// @Accept  json
// @Produce  json
// @Param accessToken query string false "Access Token"
// @Success 200 {integer} models.TokenInfo
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Authentication/Validate [get]
func (t *Transport) validate(c *gin.Context) {
	accessToken := c.Query("accessToken")

	invalid, err := t.service.IsTokenInvalid(accessToken)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if invalid {
		models.NewErrorResponse(c, http.StatusBadRequest, "token is invalid")
		return
	}

	tokenInfo, err := t.service.Validate(accessToken)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenInfo)
}

// @Summary Refresh
// @Tags Authentication
// @Description Refresh token
// @ID refresh
// @Accept  json
// @Produce  json
// @Param Input body models.RefreshRequest true "Refresh Token"
// @Success 200 {integer} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Authentication/Refresh [post]
func (t *Transport) refresh(c *gin.Context) {
	var req models.RefreshRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	invalid, err := t.service.IsTokenInvalid(req.RefreshToken)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if invalid {
		models.NewErrorResponse(c, http.StatusBadRequest, "token is invalid")
		return
	}

	token, err := t.service.Refresh(req.RefreshToken)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: token})
}
