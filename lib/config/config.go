package config

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/BurntSushi/toml"

	"go-mall/lib/config/env"
	"go-mall/lib/diagnostics"
)

var (
	prefix    = "app"
	extension = ".toml"
)

var conf *AppConfig // global

func init() {
	var once sync.Once
	once.Do(func() {
		decodeToml(configFilename(env.Env()), conf)
	})
}

// decodeToml decodes the content in toml file to struct.
// filename is the file name in root directory.
// v is pointer of struct.
func decodeToml(filename string, v interface{}) {
	fpath := path.Join(env.Root(), filename)
	if _, err := os.Stat(fpath); err != nil {
		if os.IsNotExist(err) {
			panic(fmt.Sprintf(`[config] the config file "%s" not found in root directory.`, fpath))
		} else {
			panic(err)
		}
	}

	if _, err := toml.DecodeFile(fpath, v); err != nil {
		diagnostics.WriteIf("lib_app_config_initialize", fpath, err)
		panic(err)
	}
}

// Conf get the config from the 'app.[env].conf' file in root.
// note: the config would be built when app init, and store singleton.
func Conf() *AppConfig {
	return conf
}

func configFilename(env string) string {
	if env == "" {
		return prefix + extension
	}

	return prefix + "." + env + extension
}
