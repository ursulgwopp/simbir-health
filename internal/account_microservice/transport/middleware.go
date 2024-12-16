package transport

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

func (t *Transport) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, "empty auth header")
		return
	}

	invalid, err := t.service.IsTokenInvalid(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if invalid {
		models.NewErrorResponse(c, http.StatusUnauthorized, "token is invalid")
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
		models.NewErrorResponse(c, http.StatusBadRequest, "empty auth header")
		return
	}

	invalid, err := t.service.IsTokenInvalid(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if invalid {
		models.NewErrorResponse(c, http.StatusUnauthorized, "token is invalid")
		return
	}

	tokenInfo, err := t.service.ParseToken(header)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !tokenInfo.IsAdmin {
		models.NewErrorResponse(c, http.StatusForbidden, "access denied")
		return
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("user_id")
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func getToken(c *gin.Context) (string, error) {
	id, ok := c.Get("token")
	if !ok {
		return "", errors.New("token not found")
	}

	token, ok := id.(string)
	if !ok {
		return "", errors.New("token is of invalid type")
	}

	return token, nil
}
