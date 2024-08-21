package controller

import (
	"context"
	"courseonline/middleware"
	"courseonline/models"
	"courseonline/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	storedb services.Store
}

func NewUsersController(store services.Store) *UsersController {
	return &UsersController{
		storedb: store,
	}
}

func (uc *UsersController) Signup(c *gin.Context) {
	var payload models.CreateUserReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &models.CreateUserReq{
		UserName:     payload.UserName,
		UserPassword: payload.UserPassword,
	}
	user, err := uc.storedb.Signup(c, *args)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UsersController) Sigin(c *gin.Context) {
	var payload models.CreateUserReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &models.CreateUserReq{
		UserName:     payload.UserName,
		UserPassword: payload.UserPassword,
	}
	user, err := uc.storedb.Signin(c, *args)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UsersController) Signout(c *gin.Context) {
	accessToken := c.Request.Header.Get("Authorization")

	err := uc.storedb.Signout(c, accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (uc *UsersController) GetProfile(c *gin.Context) {
	var userID string

	token := middleware.GetJWTFromHeader(c)
	//token := ""
	if token != "" {
		userID = "" //uc.serviceManager.GetIDFromToken(token)
	}

	foundUser, err := uc.storedb.FindUserByUsername(context.Background(), &userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	profile := &models.UserResponse{
		UserID:    foundUser.UserID,
		UserName:  foundUser.UserName,
		UserPhone: foundUser.UserPhone,
	}

	c.JSON(http.StatusOK, profile)
}
