package orm

import (
	"strings"
	"time"

	"go-mall/lib/log"

	"github.com/jinzhu/gorm"
)

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
	db := newDB(c, func(err error) {
		log.Errorf("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	})
	db.SetLogger(ormLog{})
	db.SingularTable(true)

	return db
}

func newDB(c *Config, errHandle func(err error)) *gorm.DB {
	db, err := gorm.Open(c.Dialect, c.DSN) // 数据库连接前的初始化，
	if err != nil {
		errHandle(err)
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

	return db
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	log.Infof(strings.Repeat("%v ", len(v)), v...)
}
