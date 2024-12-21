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


