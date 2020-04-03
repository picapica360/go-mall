package model

import "time"

// User user
type User struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// UserParam user input param
type UserParam struct {
	ID int64 `form:"id"`
}

// TableName return table name
func (*User) TableName() string {
	return tablePrefix + "user"
}
