package route

import (
	"github.com/gin-gonic/gin"
)

// Init the `oms` route.
func Init(eng *gin.Engine) {
	router := eng.Group("/admin/sms")
	{
		router.POST("/", nil)
	}
}
