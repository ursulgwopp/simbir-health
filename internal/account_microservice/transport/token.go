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

func (t *Transport) refresh(c *gin.Context) {

}
