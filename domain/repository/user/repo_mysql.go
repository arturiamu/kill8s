package user

import (
	"context"
	"gorm.io/gorm"
	"kill8s/domain/entity"

	"kill8s/infrastructure/persistence"
)

var _ Repository = new(RepoMysql)

type RepoMysql struct {
	DB *gorm.DB
}

func NewRepoMysql() *RepoMysql {
	return &RepoMysql{
		DB: persistence.GetMysql(),
	}
}

func (repo *RepoMysql) Register(ctx context.Context, user *entity.User) error {
	return repo.DB.Create(user).Error
}

func (repo *RepoMysql) Login(ctx context.Context, user *entity.User) error {
	return repo.DB.Create(user).Error
}
