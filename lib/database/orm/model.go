package orm

import "time"

// Model base
type Model struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	IsDeleted bool       `json:"is_deleted"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"` // soft delete
}
