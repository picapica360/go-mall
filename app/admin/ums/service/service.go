package service

import (
	"context"
	"sync"

	"go-mall/app/admin/ums/model"
	"go-mall/app/admin/ums/repository"

	"github.com/jinzhu/gorm"
)

// svcImpl ums service. the implement for Service.
type svcImpl struct {
	repo repository.Repository

	mutex sync.RWMutex
}

var _ Service = &svcImpl{} // compile check

// Config 服务配置
type Config struct {
	DB *gorm.DB
}

// New a service
func New(c *Config) (s Service) {
	s = &svcImpl{
		repo: repository.New(c.DB),
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
	Books(c context.Context) (books []*model.Book, err error) // only for test
}
