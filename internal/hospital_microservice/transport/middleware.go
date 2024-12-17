package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/hospital_microservice/models"
)

const acccountMicroserviceHost = "http://localhost:8081"

func parseId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return -1, err
	}

	return id, nil
}

func SendRequest(method string, url string, body io.Reader) ([]byte, int, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, -1, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}

	return data, resp.StatusCode, nil
}

func (t *Transport) userIdentity(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrEmptyAuthHeader.Error())
		return
	}

	_, code, err := SendRequest("GET", acccountMicroserviceHost+fmt.Sprintf("/api/Authentication/Validate?accessToken=%s", token), nil)
	if err != nil {
		models.NewErrorResponse(c, http.StatusForbidden, custom_errors.ErrAccessDenied.Error())
		return
	}

	if code != 200 {
		models.NewErrorResponse(c, http.StatusForbidden, custom_errors.ErrAccessDenied.Error())
		return
	}
}

func (t *Transport) adminIdentity(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrEmptyAuthHeader.Error())
		return
	}

	resp, code, err := SendRequest("GET", acccountMicroserviceHost+fmt.Sprintf("/api/Authentication/Validate?accessToken=%s", token), nil)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, custom_errors.ErrAccessDenied.Error())
		return
	}

	if code != 200 {
		models.NewErrorResponse(c, http.StatusUnauthorized, custom_errors.ErrAccessDenied.Error())
		return
	}

	var tokenClaims models.TokenClaims
	if err := json.Unmarshal(resp, &tokenClaims); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !tokenClaims.IsAdmin {
		models.NewErrorResponse(c, http.StatusForbidden, custom_errors.ErrAccessDenied.Error())
		return
	}
}
