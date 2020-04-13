package hosting

import (
	"flag"

	"go-mall/lib/config/env"
	"go-mall/lib/database/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Host host
type Host struct {
	Env string

	db     *gorm.DB
	engine *gin.Engine
}

var svcEnv = flag.String("env", "", `set the service rumtime environment, like 'development','test','production'`)

func (h *Host) AddDatabase() *Host {
	db := orm.NewDB(nil)
	defer func() {
		if db != nil {
			defer db.Close()
		}
	}()

	return h
}

func (h *Host) AddWebHost() *Host {
	return h
}

func (h *Host) Run() error {
	env.SetEnv(*svcEnv) // override
	flag.Parse()

	return nil
}
