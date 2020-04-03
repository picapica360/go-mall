package repository

import (
	"github.com/jinzhu/gorm"
)

// Repository 仓储
type Repository struct {
	DB *gorm.DB // read and write database.
}

// New 创建一个新的仓储
func New(db *gorm.DB) (repo *Repository) {
	repo = &Repository{
		DB: db,
	}
	return
}
