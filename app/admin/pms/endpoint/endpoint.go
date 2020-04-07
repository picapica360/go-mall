package endpoint

import (
	"github.com/gin-gonic/gin"
)

// Init the `oms` route.
func Init(eng *gin.Engine) {
	router := eng.Group("/admin/pms")
	{
		router.POST("/", nil)
	}
}
