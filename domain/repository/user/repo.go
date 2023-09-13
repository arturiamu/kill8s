package user

import (
	"context"
	"kill8s/domain/entity"
)

var globalRepository Repository

type Repository interface {
	Register(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, user *entity.User) error
}

func GetRepository() Repository {
	return globalRepository
}

func SetRepository(repo Repository) {
	globalRepository = repo
}
