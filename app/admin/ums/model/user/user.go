package user

import "time"

// User user
type User struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

// Param user input param
type Param struct {
	ID int64 `form:"id"`
}

// TableName return table name
func (*User) TableName() string {
	return "user"
}
