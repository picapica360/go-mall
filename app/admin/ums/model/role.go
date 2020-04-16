package model

import "time"

// Role 后台用户角色
type Role struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`       // 角色
	Desc      string    `json:"desc"`       // 角色描述
	Enabled   bool      `json:"enabled"`    // 启用状态：0->禁用；1->启用'
	Sort      int       `json:"sort"`       // 排序
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// TableName return table name
func (*Role) TableName() string {
	return tablePrefix + "role"
}

// RoleParam role input param
type RoleParam struct {
	ID int `form:"id"`
}
