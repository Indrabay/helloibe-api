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
			code = http.StatusUnauthorized
		}

		splitAuth := strings.Split(auth, " ")
		if len(splitAuth) < 2 {
			code = http.StatusUnauthorized
			return
		}

		if splitAuth[0] != "Bearer" {
			code = http.StatusUnauthorized
			return
		}

		jwtHelper := utils.NewJWT(utils.Config.SigningKey)
		claims, err := jwtHelper.ValidateToken(splitAuth[1])
		if err != nil {
			code = http.StatusUnauthorized
		}

		if code != 0 {
			c.JSON(code, gin.H{
				"code":    code,
				"message": "unauthorized user",
			})
			c.Abort()
			return
		}

		c.Set("token", splitAuth[1])
		c.Set("username", claims.Username)
		c.Set("name", claims.Name)

		c.Next()
	}
}
