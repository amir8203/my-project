package handlers

import (
	"my-project/src/config"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewUserHandler(cfg *config.Config) *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) TestFunc(c *gin.Context){
	c.JSON(200, "test")
}