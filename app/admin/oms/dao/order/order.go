package order

import (
	"context"
	"time"

	"mall/app/admin/oms/model/order"

	"mall/lib/database/orm"
	"mall/lib/log"

	"github.com/jinzhu/gorm"
)

const (
	_insertSQL = ``
	_insertSQL = ``
)

type Dao struct {
	db *gorm.DB
}

// New a order dao
func New() *Dao {
	d := &Dao{
		db: orm.NewMySQL(nil),
	}

	return d
}

func (d *Dao) Insert(ctx context.Context, m *Order, now time.Time) (err error) {
	if err = d.db.Exec(_insertSQL, now).Error; err != nil {
		log.Errorf("d.db.Exec error(%v)", err)
		return
	}
	return
}

func (m *Order) Delete() {

}

// Close close connect
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
