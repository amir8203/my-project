package middlewares

import (
	"errors"
	"fmt"
	"my-project/src/api/helper"
	"my-project/src/config"
	"my-project/src/constants"
	"my-project/src/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	tokenService := services.NewTokenService(cfg)

	return func (c *gin.Context)  {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader(constants.AuthorizationHeaderKey)
		// bearer token.token.token
		token := strings.Split(auth, " ")
		if auth == "" {
			err = fmt.Errorf("token required") 
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch  {
					case errors.Is(err, jwt.ErrTokenExpired):
						err = fmt.Errorf("token expire")
					default:
						err = fmt.Errorf("invalid token")
				}
			}
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
				nil, false, -2, err))
		}

		c.Set("userId", claimMap["user_id"])
		c.Set("exp", claimMap["exp"])
		c.Set("username", claimMap["username"])
		c.Set("phone", claimMap["phone"])

		c.Next()


	}
}