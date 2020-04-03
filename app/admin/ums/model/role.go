package model

import "time"

// Role role
type Role struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// RoleParam role input param
type RoleParam struct {
	ID int64 `form:"id"`
}

// TableName return table name
func (*Role) TableName() string {
	return tablePrefix + "role"
}
