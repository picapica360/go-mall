package service

import (
	"context"
	"sync"

	"go-mall/app/admin/ums/repository"
	"go-mall/lib/database/orm"
)

// svcImpl ums service. the implement for Service.
type svcImpl struct {
	repo repository.Repository

	mutex sync.RWMutex
}

// New a service
func New(c *orm.Config) (s *Service) {
	db := orm.NewDB(c)
	s = &svcImpl{
		repo: repository.New(db),
	}
	return
}

// Ping ping
func (s *svcImpl) Ping(c context.Context) (err error) {
	err = s.repo.Ping(c)
	return
}

// Close 关闭并释放所有占用的资源
func (s *svcImpl) Close() {
	s.repo.Close()
}

// Service interface
type Service interface {
	Ping(context.Context) error
	Close()

	// 业务
}
