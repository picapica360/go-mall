package order

import (
	"go-mall/lib/database/orm"
)

// Order order
type Order struct {
	orm.Model
}

// Param order input param
type Param struct {
	ID int64 `form:"id"`
}

// TableName return table name
func (*Order) TableName() string {
	return "oms_order"
}
