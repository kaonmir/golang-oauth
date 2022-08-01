package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")

		key := "admin"
		secret := "1234"

		if key != reqKey || secret != reqSecret {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
