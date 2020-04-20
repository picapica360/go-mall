package repository

import (
	"context"

	"go-mall/app/admin/ums/model"
)

const (
	_adminInsertSQL = ``
	_adminUpdateSQL = ``
)

// Admin 获取管理员信息
func (repo *repoImpl) Admin(c context.Context, username string) (member *model.Admin, err error) {
	err = repo.db.Where("username = ?", username).First(member).Error
	return
}

// AdminInsert admin insert
func (repo *repoImpl) AdminInsert(c context.Context, admin *model.Admin) (err error) {
	return
}
