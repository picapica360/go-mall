package main

import (
	"flag"
	"fmt"

	"net/http"

	cmsendpoint "go-mall/app/admin/cms/endpoint"
	omsendpoint "go-mall/app/admin/oms/endpoint"
	pmsendpoint "go-mall/app/admin/pms/endpoint"
	smsendpoint "go-mall/app/admin/sms/endpoint"
	umsendpoint "go-mall/app/admin/ums/endpoint"

	"go-mall/lib/config"
	"go-mall/lib/config/env"
	"go-mall/lib/database/orm"
	"go-mall/lib/log"
	httpd "go-mall/lib/net/http"

	_ "net/http/pprof"

	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var svcEnv string
	flag.StringVar(&svcEnv, "env", "", `set the service rumtime environment, like 'development','test','production'`)
	flag.Parse()

	// init config
	env.SetEnv(svcEnv) // 覆盖环境变量中的值
	config.Init()
	conf := config.Conf()
	fmt.Printf("config %+v\n", conf)

	// only listening for pprof
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%d", conf.App.PProfPort), nil)
	}()

	// init log
	log.Register(log.NewConsoleAdapter())
	log.Build()

	// service start
	fmt.Print("Server starting ...")

	// init db context.
	db := orm.NewDB(&conf.Database)
	defer func() {
		if db != nil {
			defer db.Close()
		}
	}()

	engine := httpd.Default() // middleware Logger, Recovery
	omsendpoint.Init(engine)
	pmsendpoint.Init(engine)
	cmsendpoint.Init(engine)
	smsendpoint.Init(engine)

	umsendpoint.Init(&umsendpoint.Config{
		DB:     db,
		Engine: engine,
	})

	log.Panic(engine.Run(fmt.Sprintf(":%d", conf.App.Port))) // log if starting error.
}
