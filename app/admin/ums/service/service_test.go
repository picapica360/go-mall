package service

import (
	"testing"

	"go-mall/lib/database/orm"
	"go-mall/lib/log"

	_ "github.com/go-sql-driver/mysql"
)

func testNewService() Service {
	log.Register(log.NewConsoleAdapter())
	log.Build()

	ormConf := orm.Config{
		Dialect: "mysql",
		DSN:     "mall:123456@(localhost)/mall?charset=utf8&parseTime=True&loc=Local",
	}

	db := orm.NewDB(&ormConf)
	svc := New(&Config{
		DB: db,
	})

	return svc
}

func TestSvcMember(t *testing.T) {
	svc := testNewService()
	defer svc.Close()
}
