package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaonmir/OAuth/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		env := config.Env()
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")

		var key string
		var secret string
		if key = env.AuthKey; len(strings.TrimSpace(key)) == 0 {
			c.AbortWithStatus(500)
		}
		if secret = env.AuthSecret; len(strings.TrimSpace(secret)) == 0 {
			c.AbortWithStatus(401)
		}
		if key != reqKey || secret != reqSecret {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
