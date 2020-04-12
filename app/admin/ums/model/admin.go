package model

import "time"

// Admin 后台管理员
type Admin struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`   // 用户名
	Password  string    `json:"password"`   // 密码
	Nickname  string    `json:"nickname"`   // 昵称
	Icon      string    `json:"icon"`       // 头像
	Note      string    `json:"note"`       // 备注
	Status    bool      `json:"status"`     // 帐号启用状态:0->禁用；1->启用
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	LoginAt   time.Time `json:"login_at"`   // 最后登录时间
}

// AdminInParam admin input param
type AdminInParam struct {
	ID int64 `form:"id"`
}

// TableName return table name
func (*Admin) TableName() string {
	return tablePrefix + "admin"
}
