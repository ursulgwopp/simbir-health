package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
)

func (t *Transport) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrEmptyAuthHeader.Error())
		return
	}

	invalid, err := t.service.IsTokenInvalid(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if invalid {
		models.NewErrorResponse(c, http.StatusUnauthorized, custom_errors.ErrInvalidToken.Error())
		return
	}

	tokenInfo, err := t.service.ParseToken(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("token", header)
	c.Set("user_id", tokenInfo.UserId)
	c.Set("is_admin", tokenInfo.IsAdmin)
}

func (t *Transport) adminIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrEmptyAuthHeader.Error())
		return
	}

	invalid, err := t.service.IsTokenInvalid(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if invalid {
		models.NewErrorResponse(c, http.StatusUnauthorized, custom_errors.ErrInvalidToken.Error())
		return
	}

	tokenInfo, err := t.service.ParseToken(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !tokenInfo.IsAdmin {
		models.NewErrorResponse(c, http.StatusForbidden, custom_errors.ErrAccessDenied.Error())
		return
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("user_id")
	if !ok {
		return 0, custom_errors.ErrIdNotFound
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, custom_errors.ErrInvalidUserId
	}

	return idInt, nil
}

func getToken(c *gin.Context) (string, error) {
	token_, ok := c.Get("token")
	if !ok {
		return "", custom_errors.ErrIdNotFound
	}

	token, ok := token_.(string)
	if !ok {
		return "", custom_errors.ErrInvalidTokenType
	}

	return token, nil
}
