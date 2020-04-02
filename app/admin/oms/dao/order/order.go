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
	gorm.Model
}

// New a order dao
func New() *Dao {
	d := &Dao{
		db: orm.NewMySQL(nil),
	}
	return d
}

func (d *Dao) Get(id int64) *order.Order {
	var order order.Order
	d.db.First(&order, id)
	return &order
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
func (d *Dao) Delete(order *order.Order) {
	d.Delete(order)
}

// Close close connect
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
