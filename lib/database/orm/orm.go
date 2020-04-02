package orm

import (
	"fmt"
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

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *gorm.DB) {
	c.Dialect = "mysql"
	return NewDB(c)
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

// deleteCallback used to delete data from database or set deleted_at to current time (when using with soft delete)
// ref: https://github.com/jinzhu/gorm/blob/master/callback_delete.go
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		if !scope.Search.Unscoped {
			if isDeletedField, ok := scope.FieldByName("IsDeleted"); ok {
				if deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt"); hasDeletedAtField {
					scope.Raw(fmt.Sprintf(
						"UPDATE %v SET %v=%v,%v=%v%v%v",
						scope.QuotedTableName(),
						scope.Quote(isDeletedField.DBName),
						scope.AddToVars(1),
						scope.Quote(deletedAtField.DBName),
						scope.AddToVars(gorm.NowFunc()),
						addExtraSpaceIfExist(scope.CombinedConditionSql()),
						addExtraSpaceIfExist(extraOption),
					)).Exec()
				} else {
					scope.Raw(fmt.Sprintf(
						"UPDATE %v SET %v=%v%v%v",
						scope.QuotedTableName(),
						scope.Quote(isDeletedField.DBName),
						scope.AddToVars(1),
						addExtraSpaceIfExist(scope.CombinedConditionSql()),
						addExtraSpaceIfExist(extraOption),
					)).Exec()
				}

				return
			}
		}

		scope.Raw(fmt.Sprintf(
			"DELETE FROM %v%v%v",
			scope.QuotedTableName(),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)).Exec()
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	log.Infof(strings.Repeat("%v ", len(v)), v...)
}
