package transport

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/models"
)

const host = "localhost:8082"

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
	accessToken := c.Query("accessToken")
	if accessToken == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, custom_errors.ErrEmptyAuthHeader.Error())
		return
	}

	_, code, err := SendRequest("GET", host+fmt.Sprintf("/api/Authentication/Validate?accessToken=%s", accessToken), nil)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, custom_errors.ErrAccessDenied.Error())
		return
	}

	if code != 200 {
		models.NewErrorResponse(c, http.StatusUnauthorized, custom_errors.ErrAccessDenied.Error())
		return
	}
}
