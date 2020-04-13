package main

import (
	// "flag"
	"fmt"

	// "net/http"

	cmsendpoint "go-mall/app/admin/cms/endpoint"
	omsendpoint "go-mall/app/admin/oms/endpoint"
	pmsendpoint "go-mall/app/admin/pms/endpoint"
	smsendpoint "go-mall/app/admin/sms/endpoint"
	umsendpoint "go-mall/app/admin/ums/endpoint"

	"go-mall/lib/hosting"
	"go-mall/lib/log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "net/http/pprof"
	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// service start
	fmt.Print("Server starting ...\n")

	err := hosting.WebHost.AddConsoleLogger().AddPProf().AddDB().AddRouter(func(engine *gin.Engine, db *gorm.DB) {
		omsendpoint.Init(engine)
		pmsendpoint.Init(engine)
		cmsendpoint.Init(engine)
		smsendpoint.Init(engine)
		umsendpoint.Init(&umsendpoint.Config{
			DB:     db,
			Engine: engine,
		})
	}).Run()

	log.Panic(err)
}
