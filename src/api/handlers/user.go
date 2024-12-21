package handlers

import (
	"fmt"
	"my-project/src/api/dto"
	"my-project/src/api/helper"
	"my-project/src/config"
	"my-project/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	service := services.NewUserService(cfg)
	return &UserHandler{
		service: service,
	}
}



// @Summary Login by Username
// @Description Login user by username and return a token.
// @Tags AuthenticationUsers
// @Accept json
// @Produce json
// @Param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 500 {object} helper.BaseHttpResponse "other error"
// @Router /v1/users/login [post]
func (h *UserHandler) LoginByUsername(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	token, err := h.service.LoginByUsername(req.Username, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(token, true, 0))
}



// @Summary register by Username
// @Description register user by username
// @Tags AuthenticationUsers
// @Accept json
// @Produce json
// @Param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 500 {object} helper.BaseHttpResponse "other error"
// @Router /v1/users/register [post]
func (h *UserHandler) RegisterByUsername(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.RegisterByUsername(*req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}



// @Summary Show Info
// @Description show info by token
// @Tags UsersProfile
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/profile [get]
// @security Bearer
func (h *UserHandler) ShowProfile(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponseWithError(nil, false, -1, fmt.Errorf("user id not found")))
		return
	}

	user, err := h.service.GetInfo(int(userId.(float64)))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponseWithError(nil, false, -1, fmt.Errorf("user id not found")))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(user, true, 0))
}


// @Summary Update information
// @Description Update Information by token and params (all params are optional)
// @Tags UsersProfile
// @Accept  json
// @Produce  json
// @Param Request body dto.UpdateUserProfileRequest true "UpdateUserProfileRequest"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/profile [put]
// @security Bearer
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	req := new(dto.UpdateUserProfileRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponseWithError(nil, false, -1, fmt.Errorf("user id not found")))
		return
	}

	err = h.service.UpdateUserProfile(*req ,int(userId.(float64)))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("update is successful", true, 0))

}


// @Summary delete account
// @Description delete account by token
// @Tags UsersProfile
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 500 {object} helper.BaseHttpResponse "Failed"
// @Failure 404 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/profile [delete]
// @security Bearer
func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponseWithError(nil, false, -1, fmt.Errorf("user id not found")))
		return
	}

	err := h.service.DeleteAccount(int(userId.(float64)))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			helper.GenerateBaseResponseWithError(nil, false, -1, fmt.Errorf("cant delete account")))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("delete account is successful", true, 0))

}


