package repository

import (
	"context"

	"go-mall/app/admin/ums/model"
)

const (
	_memberInsertSQL = ``
	_memberUpdateSQL = ``
)

func (repo *repoImpl) Books(c context.Context) (books []*model.Book, err error) {
	err = repo.db.Find(books).Error
	return
}

// GetUserByID inherit: Repository.GetUserByID
func (repo *repoImpl) GetMemberByID(c context.Context, id int64) (member *model.Member, err error) {
	err = repo.db.First(member, id).Error
	return
}

// User 获取用户信息
func (repo *repoImpl) Member(c context.Context, username string) (member *model.Member, err error) {
	err = repo.db.Where("username = ?", username).First(member).Error
	return
}
