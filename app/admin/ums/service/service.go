package service

import (
	"context"
	"sync"

	"go-mall/app/admin/ums/repository"
	"go-mall/lib/database/orm"

	"github.com/jinzhu/gorm"
)

// Service ums service.
type Service struct {
	repo repository.Repository

	db    *gorm.DB
	mutex sync.RWMutex
}

// New a service
func New(c *orm.Config) s *Service) {
	db := orm.NewDB()
	s = &Service{
		repo: repository.New(db),
		db:   db,
	}
	return
}

// Ping Service
func (s *Service) Ping(c context.Context) (err error) {
	err = s.db.DB().PingContext(c)
	return
}

// Close the service
func (s *Service) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
