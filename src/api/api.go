package api

import (
	"fmt"
	"my-project/src/config"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}