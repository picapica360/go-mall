package model

import "time"

// Resource 后台资源
type Resource struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Name       string    `json:"name"`        // 资源名称
	URI        string    `json:"uri"`         // 资源URI
	Desc       string    `json:"desc"`        // 描述
	CategoryID int       `json:"category_id"` // 资源分类 id, ref: ResourceCategory
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
}

// TableName return table name
func (*Resource) TableName() string {
	return tablePrefix + "resource"
}

// ResourceCategory 后台资源分类
type ResourceCategory struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      int       `json:"name"`       // 资源分类名称
	Sort      int       `json:"sort"`       // 排序
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// TableName return table name
func (*ResourceCategory) TableName() string {
	return tablePrefix + "resource_category"
}
