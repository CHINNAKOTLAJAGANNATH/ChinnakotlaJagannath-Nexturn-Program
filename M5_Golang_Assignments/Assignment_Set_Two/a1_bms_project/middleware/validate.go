package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() != "application/json" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Content type must be application/json"})
			c.Abort()
			return
		}
		c.Next()
	}
}
