package handlers

import (
	"fmt"
	"my-project/src/api/dto"
	"my-project/src/api/helper"
	"my-project/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	// UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	return &UserHandler{
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
	
	//call service
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("response", true, 0))
}



func (h *UserHandler) RegisterByUsername(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	
	//call service

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
}




func (h *UserHandler) ShowProfile(c *gin.Context) {
	//get user id from jwt then from context
	userId, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponseWithError(nil, false, -1, fmt.Errorf("user id= %d not found", userId)))
		return
	}

	//call service

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("user", true, 0))
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
			helper.GenerateBaseResponseWithError(nil, false, -1, fmt.Errorf("user id= %d not found", userId)))
		return
	}


	//call service

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("update is successful", true, 0))

}



func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponseWithError(nil, false, -1, fmt.Errorf("user id= %d not found", userId)))
		return
	}

	//call service

	c.JSON(http.StatusOK, helper.GenerateBaseResponse("delete account is successful", true, 0))

}


