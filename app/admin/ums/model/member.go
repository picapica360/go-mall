package model

import (
	"time"
)

// Member 会员
type Member struct {
	ID                    int64     `gorm:"primary_key" json:"id"`
	Username              string    `json:"username"`
	Password              string    `json:"password"`
	Nickname              string    `json:"nickname"`
	Phone                 string    `json:"phone"`
	Status                bool      `json:"status"`
	Icon                  string    `json:"icon"`
	Gender                int       `json:"gender"` // 性别：0->未知；1->男；2->女
	Birthday              time.Time `json:"birthday"`
	City                  string    `json:"city"`
	Job                   string    `json:"job"`
	PersonalizedSignature string    `gorm:"column:personalized_signature" json:"personalized_signature"`
	SourceType            string    `gorm:"column:source_type" json:"source_type"`
	Integration           int       `json:"integration"`
	Growth                int       `json:"growth"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// MemberParam member input param
type MemberParam struct {
}

// TableName return table name
func (*Member) TableName() string {
	return tablePrefix + "user"
}
