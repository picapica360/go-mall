package service

import (
	"context"

	"go-mall/app/admin/ums/model"
)

func (s *svcImpl) Books(c context.Context) (books []*model.Book, err error) {
	return s.repo.Books(c)
}

// Member 获取会员相关信息
func (s *svcImpl) Member(username string) (member *model.Member) {
	return
}
