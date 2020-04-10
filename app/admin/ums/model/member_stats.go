package model

// MemberStats 会员统计信息
type MemberStats struct {
	ID int64 `gorm:"primary_key" json:"id"`
}

// TableName return table name
func (*MemberStats) TableName() string {
	return tablePrefix + "admin_stats"
}
