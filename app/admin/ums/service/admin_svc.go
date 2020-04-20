package service

import (
	"context"
	"strings"

	"go-mall/lib/utils"

	"go-mall/app/admin/ums/model"
)

// GetRoles 获取用户的全部角色
func (s *svcImpl) GetRoles(username string) (roles []*model.Role) {
	return
}

func (s *svcImpl) Register(username string) (roles []*model.Role) {
	return
}

// CkeckAdmin 检查用户姓名和密码
func (s *svcImpl) CkeckAdmin(c context.Context, username, password string) (ret bool, err error) {
	// 从仓储中获取用户密码
	var pwd string
	m5 := utils.MD5String(password)
	ret = strings.EqualFold(m5, pwd)

	return
}
