package orm

import (
	"strings"
	"time"

	"go-mall/lib/log"

	// MySql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Config mysql config.
type Config struct {
	Dialect     string        // db name, like "mysql", "mssql", "postgres" or "sqlite"
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}

func init() {
	// replace gorm delete with custom callback
	gorm.DefaultCallback.Delete().Replace("gorm:delete", deleteCallback)
}

// NewDB create a new db connection.
// Firstly, must import driver in main().
// eg:
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	_ "github.com/jinzhu/gorm/dialects/mssql"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
func NewDB(c *Config) *gorm.DB {
	db, err := gorm.Open(c.Dialect, c.DSN)
	if err != nil {
		log.Errorf("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}

	if c.Idle > 0 {
		db.DB().SetMaxIdleConns(c.Idle)
	}
	if c.Active > 0 {
		db.DB().SetMaxOpenConns(c.Active)
	}
	if c.IdleTimeout > time.Duration(0) {
		db.DB().SetConnMaxLifetime(c.IdleTimeout)
	}

	db.SetLogger(ormLog{})

	return db
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	log.Infof(strings.Repeat("%v ", len(v)), v...)
}
