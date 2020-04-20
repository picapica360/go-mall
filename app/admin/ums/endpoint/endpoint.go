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
	r1 := engine.Group("/admin")
	{
		r1.POST("/login", ctl.AdminLogin)
		r1.GET("/info", nil)
		r1.POST("/logout", ctl.AdminLogout)
		r1.GET("/list", nil)
		r1.POST("/register", ctl.AdminRegister)
		r1.POST("/update/:id", nil)
		r1.POST("/delete/:id", nil)

		r1.GET("/role/:id", ctl.AdminRole)
		r1.POST("/role/update", nil)
	}
}
