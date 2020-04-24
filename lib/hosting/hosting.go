package hosting

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	C Context // Context

	conf *config.AppConfig

	middlewareFn []func() gin.HandlerFunc
	logFn        func()
	endpointFn   func(Context)
	dbFn         func() *gorm.DB
	healthFn     func(*gin.Engine)
	pprofFn      func()
}

// Context Host Context
type Context struct {
	Engine *gin.Engine
	DB     *gorm.DB
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

// AddEndpoint add web endpoint.
func (h *Host) AddEndpoint(fn func(c Context)) *Host {
	h.endpointFn = fn
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

// AddHealth add health check api.
// note: check the database, redis, ES etc.
func (h *Host) AddHealth() *Host {
	h.healthFn = func(engine *gin.Engine) {
		engine.GET("/health", func(c *gin.Context) {
			type errModel struct {
				Name string
				Err  error
			}
			var errs []errModel
			if h.C.DB != nil {
				if err1 := h.C.DB.DB().PingContext(context.TODO()); err1 != nil {
					errs = append(errs, errModel{"database", err1})
				}
			}

			if len(errs) > 0 {
				c.JSON(http.StatusBadRequest, errs)
			} else {
				c.JSON(http.StatusOK, nil)
			}
		})
	}
	return h
}

// Run run the hosting
// note: env->config->pprof->log->dbcontext->health->webhost
func (h *Host) Run() {
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
		h.C.DB = h.dbFn()
		// close the database
		defer func() {
			if h.C.DB != nil {
				fmt.Println("database closed")
				h.C.DB.Close()
			}
		}()
	}

	// web host
	h.C.Engine = httpd.Default()

	if len(h.middlewareFn) > 0 {
		for _, fn := range h.middlewareFn {
			h.C.Engine.Use(fn())
		}
	}

	if h.endpointFn != nil {
		h.endpointFn(h.C)
	}
	port := h.conf.App.Port
	if port == 0 {
		port = defaultHostPort
	}

	if h.healthFn != nil {
		h.healthFn(h.C.Engine)
	}

	// h.C.Engine.Run(fmt.Sprintf(":%d", port))
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      h.C.Engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// listen serve
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Infof("listen: %v", err)
		}
	}()

	shutdown(srv)
}

// // wait for interrupt signal to close server (timeout: 5s)
func shutdown(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Infof("get a signal %s, stop the process", sig.String())

	fmt.Println("shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		// note: do not call Fatal, because it will call os.Exit(),
		// 	and the 'defer func' (include caller) will not be executed.
		log.Infof("server shutdown error: %v", err)
	}
	fmt.Println("server exited")
}
