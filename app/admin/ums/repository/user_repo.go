package repository

import (
	"context"

	"go-mall/app/admin/ums/model"
)

const (
	_insertSQL = ``
	_updateSQL = ``
)

// GetUserByID inherit: Repository.GetUserByID
func (repo *repoImpl) GetUserByID(c context.Context, id int64) (user *model.User, err error) {
	err = repo.db.First(user, id).Error
	return
}

// User 获取用户信息
func (repo *repoImpl) User(c context.Context, username string) (user *model.User, err error) {
	err = repo.db.Where("username = ?", username).First(user).Error
	return
}
