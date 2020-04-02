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
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		log.Errorf("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}
	if c.Active > 0 {
		db.DB().SetMaxIdleConns(c.Idle)
		db.DB().SetMaxOpenConns(c.Active)
	}
	if c.Idle > 0 {
		db.DB().SetMaxIdleConns(c.Idle)
	}
	if c.IdleTimeout > time.Duration(0) {
		db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	}
	db.SetLogger(ormLog{})
	return
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	log.Infof(strings.Repeat("%v ", len(v)), v...)
}
