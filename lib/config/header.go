package config

import (
	"go-mall/lib/database/orm"
	"go-mall/lib/log"
)

// AppConfig app config, mapping 'app.[env].toml' file.
type AppConfig struct {
	App struct {
		Port      int // main service listening port
		PProfPort int // pprof listening port
	}
	Version  string // the version of app.
	Log      log.Config
	Database orm.Config
}
