package order

import (
	"context"
	"time"

	"mall/lib/database/orm"

	"github.com/jinzhu/gorm"
)

const (
	_insertSQL := ``
)

type Dao struct {
	db *gorm.DB
}

// Order 订单模型
type Order struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// New a order dao
func New() *Dao {
	d = &Dao{
		db: orm.NewMySQL(nil),
	}

	return d
}

func (d *Dao) Insert(ctx context.Context, order *Order, now time.Time) (err error) {
	
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
