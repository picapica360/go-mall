package model

import "time"

// Permission 后台用户权限
type Permission struct {
	ID        int       `gorm:"primary_key" json:"id"`
	PID       int       `json:"pid"`        // 父级权限id, 0 表示最top
	Name      string    `json:"name"`       // 名称
	Value     string    `json:"value"`      // 权限值
	Icon      string    `json:"icon"`       // 图标
	Kind      int       `json:"kind"`       // 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）
	URI       string    `json:"uri"`        // 前端资源路径
	Enabled   bool      `json:"enabled"`    // 启用状态：0->禁用；1->启用'
	Sort      int       `json:"sort"`       // 排序
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

// PermissionInParam permission input param
type PermissionInParam struct {
	ID int `form:"id"`
}

// TableName return table name
func (*Permission) TableName() string {
	return tablePrefix + "permission"
}
