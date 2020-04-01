package order

import (
	"time"
)

// Order order
type Order struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// Param param
type Param struct {
	ID       int64  `form:"id"`
}

// TableName return table name
func (*Order) TableName() string {
	return "order"
}
