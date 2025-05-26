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
			unauthorized(c)
			return
		}

		splitAuth := strings.Split(auth, " ")
		if len(splitAuth) < 2 {
			unauthorized(c)
			return
		}

		if splitAuth[0] != "Bearer" {
			unauthorized(c)
			return
		}

		jwtHelper := utils.NewJWT(utils.Config.SigningKey)
		claims, err := jwtHelper.ValidateToken(splitAuth[1])
		if err != nil {
			unauthorized(c)
			return
		}

		if code != 0 {
			unauthorized(c)
			return
		}

		c.Set("token", splitAuth[1])
		c.Set("username", claims.Username)
		c.Set("name", claims.Name)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "unauthorized user",
	})
	c.Abort()
}
