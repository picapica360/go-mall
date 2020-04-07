package endpoint

import (
	"go-mall/app/admin/ums/controller"
	"go-mall/app/admin/ums/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Config endpoint config.
type Config struct {
	DB     *gorm.DB
	Engine *gin.Engine
}

// Init the `oms` route.
func Init(c *Config) {
	conf := &controller.Config{
		Svc: service.New(&service.Config{
			DB: c.DB,
		}),
	}
	ctl := controller.New(conf)
	engine := c.Engine
	router := engine.Group("/admin/ums")
	{
		router.POST("/user", ctl.GetUser) // get user info
	}
}
