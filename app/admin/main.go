package main

import (
	"net/http"

	cmsendpoint "go-mall/app/admin/cms/endpoint"
	omsendpoint "go-mall/app/admin/oms/endpoint"
	pmsendpoint "go-mall/app/admin/pms/endpoint"
	smsendpoint "go-mall/app/admin/sms/endpoint"
	umsendpoint "go-mall/app/admin/ums/endpoint"

	"go-mall/lib/config"
	"go-mall/lib/database/orm"
	"go-mall/lib/log"

	"github.com/gin-gonic/gin"

	_ "net/http/pprof"

	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// only listening for pprof
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	// init log
	log.Register(log.NewConsoleAdapter())
	log.Build()

	// init config
	conf := config.Conf()
	log.Infof("config: %v\n", conf) // only for trace, development environment.

	// service start
	log.Info("Server starting ...")

	// init db context.
	db := orm.NewDB(&conf.Database)
	defer func() {
		if db != nil {
			defer db.Close()
		}
	}()

	engine := gin.Default() // middleware Logger, Recovery
	omsendpoint.Init(engine)
	pmsendpoint.Init(engine)
	cmsendpoint.Init(engine)
	smsendpoint.Init(engine)

	umsendpoint.Init(&umsendpoint.Config{
		DB:     db,
		Engine: engine,
	})

	log.Panic(engine.Run(":5000")) // log if starting error.
}
