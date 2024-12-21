package routers

import (
	"my-project/src/api/handlers"
	"my-project/src/api/middlewares"
	"my-project/src/config"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUserHandler(cfg)

	r.POST("/login", h.LoginByUsername)
	r.POST("/register", h.RegisterByUsername)

	profile := r.Group("/profile")
	profile.Use(middlewares.Authentication(cfg))

	profile.GET("/", h.ShowProfile)
	profile.PUT("/", h.UpdateProfile)
	profile.DELETE("/", h.DeleteAccount)
}