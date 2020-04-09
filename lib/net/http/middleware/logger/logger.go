package middleware

import (
	"github.com/gin-gonic/gin"
)

// Logger http logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()
	}
}
