package endpoint

import (
	"go-mall/app/admin/ums/controller"
	"go-mall/lib/database/orm"

	"github.com/gin-gonic/gin"
)

// Init the `oms` route.
func Init(c orm.Config, engine *gin.Engine) {
	router := engine.Group("/admin/ums")
	{
		router.POST("/user", controller.GetUser) // get user info
	}
}
