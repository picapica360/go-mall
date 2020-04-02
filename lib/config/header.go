package config

import (
	"go-mall/lib/database/orm"
	"go-mall/lib/log"
)

// AppConfig app config, mapping 'app.[env].toml' file.
type AppConfig struct {
	Version  string // the version of app.
	Log      log.Config
	Database orm.Config
}
