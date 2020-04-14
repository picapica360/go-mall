package model

import "time"

// Book for test
type Book struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Author    string     `json:"author"`
	Desc      string     `json:"desc"`
	ISBN      string     `json:"ISBN"`
	Pub       time.Time  `json:"pub"`
	Amount    float64    `json:"amount"`
	IsPutaway bool       `json:"is_putaway"`
	IsDeleted bool       `json:"is_deleted"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (*Book) TableName() string {
	return "test_book"
}
