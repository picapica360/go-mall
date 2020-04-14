package repository

import (
	"context"

	"go-mall/app/admin/ums/model"

	"github.com/jinzhu/gorm"
)

// RepoImpl 仓储, 也是 Repository 接口的实现。
type repoImpl struct {
	db *gorm.DB // read and write database.
}

var _ Repository = &repoImpl{} // compile check

// New 创建一个新的仓储
func New(db *gorm.DB) (repo Repository) {
	repo = &repoImpl{
		db: db,
	}
	return
}

// Ping ping
func (repo *repoImpl) Ping(ctx context.Context) error {
	return repo.db.DB().PingContext(ctx)
}

// Close the repository.
// 何时调用 close？
// 在驱动 "database/sql" 内部会维护一个连接池，使用者不需要管连接的创建和关闭。
func (repo *repoImpl) Close() {
	if repo.db != nil {
		repo.db.DB().Close()
	}
}

// Repository 仓储接口
// 该仓储是针对每个资源库的，而不是每个实体对象的
type Repository interface {
	Ping(context.Context) error
	Close()

	// 业务
	Books(c context.Context) (books []*model.Book, err error)
	Member(c context.Context, username string) (member *model.Member, err error)
}
