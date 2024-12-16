package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

// @Router /api/Authentication/Validate [get]
// @Summary Validate
// @Tags Authentication
// @Description Validate token
// @ID validate
// @Accept json
// @Produce json
// @Param accessToken query string false "Access Token"
// @Success 200 {object} models.TokenInfo
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) validate(c *gin.Context) {
	// CHECKING IF TOKEN IS INVALID
	accessToken := c.Query("accessToken")
	if accessToken == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrEmptyAuthHeader.Error())
		return
	}

	invalid, err := t.service.IsTokenInvalid(accessToken)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if invalid {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrInvalidToken.Error())
		return
	}

	// PASSING TOKEN TO SERVICE LAYER
	tokenInfo, err := t.service.Validate(accessToken)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenInfo)
}

// @Router /api/Authentication/Refresh [post]
// @Summary Refresh
// @Tags Authentication
// @Description Refresh token
// @ID refresh
// @Accept json
// @Produce json
// @Param Input body models.RefreshRequest true "Refresh Token"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) refresh(c *gin.Context) {
	// UNMARSHALLING REQUEST BODY
	var req models.RefreshRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// CHECKING IF TOKEN IS INVALID
	invalid, err := t.service.IsTokenInvalid(req.RefreshToken)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if invalid {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrInvalidToken.Error())
		return
	}

	// PASSING BODY TO SERVICE LAYER
	token, err := t.service.Refresh(req.RefreshToken)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
