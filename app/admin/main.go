package main

import (
	"fmt"

	cmsep "go-mall/app/admin/cms/endpoint"
	omsep "go-mall/app/admin/oms/endpoint"
	pmsep "go-mall/app/admin/pms/endpoint"
	smsep "go-mall/app/admin/sms/endpoint"
	umsep "go-mall/app/admin/ums/endpoint"

	"go-mall/lib/hosting"

	_ "net/http/pprof"
	// MySql driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// service start
	fmt.Println("Server starting ...")

	host := hosting.WebHost
	host.AddConsoleLogger()
	host.AddPProf()
	host.AddDB()
	host.AddHealth()
	host.AddEndpoint(func(c hosting.Context) {
		omsep.Init(c.Engine)
		pmsep.Init(c.Engine)
		cmsep.Init(c.Engine)
		smsep.Init(c.Engine)
		umsep.Init(&umsep.Config{Engine: c.Engine, DB: c.DB})
	})

	host.UseDefaultCookieSession()

	host.Run()
}
