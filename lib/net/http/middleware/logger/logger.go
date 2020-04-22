package middleware

import (
	"github.com/gin-gonic/gin"
)

// Logger http logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: log before handle.

		c.Next()

		// TODO: log after handle.
	}
}
