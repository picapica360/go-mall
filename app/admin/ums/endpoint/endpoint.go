package endpoint

import (
	"go-mall/app/admin/ums/controller"
	"go-mall/app/admin/ums/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Config endpoint config.
type Config struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

// Init the `oms` route.
func Init(c *Config) {
	conf := &controller.Config{
		Svc: service.New(&service.Config{DB: c.DB}),
	}
	ctl := controller.New(conf)
	engine := c.Engine
	router := engine.Group("/admin/ums")
	{
		router.GET("/test/books", ctl.Books) // test
		router.GET("/member", ctl.Member)    // get user info
	}
}
