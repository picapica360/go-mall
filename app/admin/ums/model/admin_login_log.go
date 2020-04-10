package model

// AdminLoginLog 后台管理员登录日志
type AdminLoginLog struct {
	ID int64 `gorm:"primary_key" json:"id"`
}

// TableName return table name
func (*AdminLoginLog) TableName() string {
	return tablePrefix + "admin_login_log"
}
