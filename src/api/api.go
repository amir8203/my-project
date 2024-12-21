package api

import (
	"fmt"
	"my-project/src/api/routers"
	"my-project/src/config"

	validation "my-project/src/api/validations"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidators()

	r.Use(gin.Logger(), gin.Recovery())

	RegisterRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}


func RegisterValidators() {

	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			test := v1.Group("/test")
			routers.Test(test, cfg)
			users := v1.Group("/users")
			routers.User(users, cfg)
		}
	}
}

