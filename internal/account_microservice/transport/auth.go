package transport

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
	"github.com/ursulgwopp/simbir-health/internal/models"
)

// @Router /api/Authentication/SignUp [post]
// @Summary SignUp
// @Tags Authentication
// @Description Create New Account
// @ID sign-up
// @Accept json
// @Produce json
// @Param Input body models.SignUpRequest true "Sign Up Info"
// @Success 201 {object} models.Response
// @Failure 400,409 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) signUp(c *gin.Context) {
	// UNMARSHALLING REQUEST BODY
	var req models.SignUpRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// PASSING BODY TO SERVICE LAYER
	id, err := t.service.SignUp(req)
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

// @Summary SignIn
// @Tags Authentication
// @Description Sign Into Account
// @ID sign-in
// @Accept json
// @Produce json
// @Param Input body models.SignInRequest true "Sign In Info"
// @Success 200 {object} models.Response
// @Failure 400,404,409 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /api/Authentication/SignIn [post]
func (t *Transport) signIn(c *gin.Context) {
	// UNMARSHALLING REQUEST BODY
	var req models.SignInRequest

	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// PASSING BODY TO SERVICE LAYER
	token, err := t.service.SignIn(req)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUsernameInvalidCharacters) ||
			errors.Is(err, custom_errors.ErrUsernameInvalidLength) ||
			errors.Is(err, custom_errors.ErrShortPassword) ||
			errors.Is(err, custom_errors.ErrPasswordWithoutDigits) {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, custom_errors.ErrSignIn) {
			models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Router /api/Authentication/SignOut [put]
// @Summary SignOut
// @Security ApiKeyAuth
// @Tags Authentication
// @Description Sign Out from Account
// @ID sign-out
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
func (t *Transport) signOut(c *gin.Context) {
	// GETTING AUTH TOKEN
	token, err := getToken(c)
	if err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// PASSING TOKEN TO SERVICE LAYER
	if err := t.service.SignOut(token); err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "successfully signed out"})
}

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
