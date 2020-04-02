package order

import (
	"context"
	"time"

	"go-mall/app/admin/oms/model/order"

	"go-mall/lib/database/orm"
	"go-mall/lib/log"

	"github.com/jinzhu/gorm"
)

const (
	_insertSQL = `insert into oms_order(id) values(?)`
	_updateSQL = ``
)

// Dao dao
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

// Insert order
func (d *Dao) Insert(ctx context.Context, m *order.Param, now time.Time) (err error) {
	if err = d.db.Exec(_insertSQL, now).Error; err != nil {
		log.Errorf("d.db.Exec error(%v)", err)
		return
	}
	return
}

// Delete order
func (d *Dao) Delete(id int64) {

}

// Close close connect
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
