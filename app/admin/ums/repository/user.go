package repository

import (
	"context"

	"go-mall/app/admin/ums/model"
)

const (
	_insertSQL = ``
	_updateSQL = ``
)

// User 获取用户信息
func (repo *Repository) User(c context.Context, username string) (user *model.User, err error) {
	return
}
