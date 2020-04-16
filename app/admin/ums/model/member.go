package model

import (
	"time"
)

// Member 会员信息
type Member struct {
	ID          int64     `gorm:"primary_key" json:"id"`                 // id
	Username    string    `json:"username"`                              // 用户名
	Password    string    `json:"password"`                              // 密码
	Nickname    string    `json:"nickname"`                              // 昵称
	Phone       string    `json:"phone"`                                 // 手机号
	Status      bool      `json:"status"`                                // 帐号启用状态:0->禁用；1->启用
	Icon        string    `json:"icon"`                                  // 头像
	Gender      int       `json:"gender"`                                // 性别：0->未知；1->男；2->女
	Birthday    time.Time `json:"birthday"`                              // 出生日期
	City        string    `json:"city"`                                  // 所在城市
	Job         string    `json:"job"`                                   // 职业
	Signature   string    `json:"signature"`                             // 个性签名
	SourceType  string    `gorm:"column:source_type" json:"source_type"` // 用户来源
	Integration int       `json:"integration"`                           // 积分
	Growth      int       `json:"growth"`                                // 成长值
	CreatedAt   time.Time `json:"created_at"`                            // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`                            // 修改时间
}

// MemberParam member input param
type MemberParam struct {
}

// TableName return table name
func (*Member) TableName() string {
	return tablePrefix + "member"
}
