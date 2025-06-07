package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/indrabay/helloibe-api/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 0
		request := c.Request
		auth := request.Header.Get("Authorization")
		if len(auth) < 1 {
			Unauthorized(c)
			return
		}

		splitAuth := strings.Split(auth, " ")
		if len(splitAuth) < 2 {
			Unauthorized(c)
			return
		}

		if splitAuth[0] != "Bearer" {
			Unauthorized(c)
			return
		}

		jwtHelper := utils.NewJWT(utils.Config.SigningKey)
		claims, err := jwtHelper.ValidateToken(splitAuth[1])
		if err != nil {
			Unauthorized(c)
			return
		}

		if code != 0 {
			Unauthorized(c)
			return
		}

		c.Set("token", splitAuth[1])
		c.Set("username", claims.Username)
		c.Set("name", claims.Name)
		c.Set("role", claims.Role)
		c.Set("stores", claims.Stores)

		c.Next()
	}
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "unauthorized user",
	})
	c.Abort()
}
