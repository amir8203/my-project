package routers

import (
	"my-project/src/api/handlers"
	"my-project/src/config"

	"github.com/gin-gonic/gin"
)

func Test(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewTestHandler(cfg)

	r.POST("/", h.TestFunc)
}