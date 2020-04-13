package hosting

import (
	"flag"
	"fmt"
	"net/http"

	"go-mall/lib/config"
	"go-mall/lib/config/env"
	"go-mall/lib/database/orm"
	"go-mall/lib/log"
	httpd "go-mall/lib/net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Host host
type Host struct {
	conf *config.AppConfig
	db   *gorm.DB

	logFn   func()
	routeFn func(*gin.Engine)
	dbFn    func() *gorm.DB
	pprofFn func()
}

var (
	// WebHost web host
	WebHost Host

	cliEnv *string
)

const (
	defaultHostPort  = 5000
	defaultPProfPort = 6060
)

func init() {
	cliEnv = flag.String("env", "", `set the service rumtime environment, like 'development','test' or 'production'`)
	flag.Parse()
}

// AddFileLogger add logger that output file.
func (h *Host) AddFileLogger() *Host {
	h.logFn = func() {
		log.Register(log.NewFileAdapter2(h.conf.Log))
		log.Build()
	}
	return h
}

// AddConsoleLogger add logger that output console.
func (h *Host) AddConsoleLogger() *Host {
	h.logFn = func() {
		log.Register(log.NewConsoleAdapter())
		log.Build()
	}
	return h
}

// AddLogger add logger.
func (h *Host) AddLogger(adapter log.Instance) *Host {
	h.logFn = func() {
		log.Register(adapter)
		log.Build()
	}
	return h
}

// AddDBFunc add a database service.
func (h *Host) AddDBFunc(fn func(cfg *orm.Config) *orm.Config) *Host {
	h.dbFn = func() *gorm.DB {
		conf := &orm.Config{}
		c := fn(conf)
		return orm.NewDB(c)
	}
	return h
}

// AddDB add database service by config.
func (h *Host) AddDB() *Host {
	h.dbFn = func() *gorm.DB {
		return orm.NewDB(&h.conf.Database)
	}
	return h
}

// AddRouter add web router.
func (h *Host) AddRouter(fn func(*gin.Engine, *gorm.DB)) *Host {
	h.routeFn = func(engine *gin.Engine) {
		fn(engine, h.db) // TODO: 思考如何处理 DbContext 参数
	}
	return h
}

// AddPProf only listening for pprof
// note: must import _ "net/http/pprof" package.
func (h *Host) AddPProf() *Host {
	h.pprofFn = func() {
		go func() {
			port := h.conf.App.PProfPort
			if port == 0 {
				port = defaultPProfPort
			}
			http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		}()
	}
	return h
}

// Run run the hosting
func (h *Host) Run() (err error) {
	env.SetEnv(*cliEnv) // override

	// config
	config.Init()
	h.conf = config.Conf()

	// pprof
	if h.pprofFn != nil {
		h.pprofFn()
	}

	// log
	if h.logFn != nil {
		h.logFn()
	}

	// db context
	if h.dbFn != nil {
		db := h.dbFn()
		defer func() {
			if db != nil {
				defer db.Close()
			}
		}()
	}

	// web host
	engine := httpd.Default()
	if h.routeFn != nil {
		h.routeFn(engine)
	}
	port := h.conf.App.Port
	if port == 0 {
		port = defaultHostPort
	}
	return engine.Run(fmt.Sprintf(":%d", port))
}
