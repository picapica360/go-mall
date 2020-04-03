package main

import (
	"net/http"

	cmsroute "go-mall/app/admin/cms/route"
	omsroute "go-mall/app/admin/oms/route"
	pmsroute "go-mall/app/admin/pms/route"
	smsroute "go-mall/app/admin/sms/route"
	umsendpoint "go-mall/app/admin/ums/endpoint"

	"go-mall/lib/config"
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

	engine := gin.Default() // middleware Logger, Recovery
	omsroute.Init(engine)
	pmsroute.Init(engine)
	cmsroute.Init(engine)
	smsroute.Init(engine)
	umsendpoint.Init(conf.Database, engine)
	log.Panic(engine.Run(":5000")) // log if starting error.
}
