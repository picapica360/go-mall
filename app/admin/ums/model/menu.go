package model

import "time"

// Menu 后台菜单
type Menu struct {
	ID        int       `gorm:"primary_key" json:"id"`
	PID       int       `json:"pid"`        // 父菜单id, 0 表示最top
	Title     string    `json:"title"`      // 菜单名称
	Level     int       `json:"level"`      // 菜单级数
	Sort      int       `json:"sort"`       // 菜单排序
	Name      string    `json:"name"`       // 前端名称
	Icon      string    `json:"icon"`       // 前端图标
	Hidden    bool      `json:"is_hidden"`  // 前端隐藏 0->不隐藏；1->隐藏
	CreatedAt time.Time `json:"created_at"` // 创建时间
}

// TableName return table name
func (*Menu) TableName() string {
	return tablePrefix + "menu"
}
