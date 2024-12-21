package api

import (
	"fmt"
	"my-project/src/api/routers"
	"my-project/src/config"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	RegisterRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			test := v1.Group("/test")
			routers.Test(test, cfg)
		}
	}
}