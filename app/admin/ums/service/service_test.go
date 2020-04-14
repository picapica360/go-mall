package service

import (
	"context"
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

func TestSvcBook(t *testing.T) {
	svc := testNewService()
	defer svc.Close()

	books, err := svc.Books(context.TODO())
	if err != nil {
		t.Error(err)
	}
	if len(books) == 0 {
		t.Errorf("not found the book.")
	}
}
