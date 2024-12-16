package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/account_microservice/models"
)

// @Summary SignUp
// @Tags Authentication
// @Description Create New Account
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param Input body models.SignUpRequest true "Sign Up Info"
// @Success 200 {integer} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Authentication/SignUp [post]
func (t *Transport) signUp(c *gin.Context) {
	var req models.SignUpRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := t.service.SignUp(req)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, models.Response{Message: id})
}

// @Summary SignIn
// @Tags Authentication
// @Description Sign Into Account
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param Input body models.SignInRequest true "Sign In Info"
// @Success 200 {integer} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Authentication/SignIn [post]
func (t *Transport) signIn(c *gin.Context) {
	var req models.SignInRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := t.service.SignIn(req)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: token})
}

// @Summary SignOut
// @Security ApiKeyAuth
// @Tags Authentication
// @Description Sign Out from Account
// @ID sign-out
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Authentication/SignOut [put]
func (t *Transport) signOut(c *gin.Context) {
	token, err := getToken(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.service.SignOut(token); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "ok"})
}
