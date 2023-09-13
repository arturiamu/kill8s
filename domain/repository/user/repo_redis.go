package user

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"kill8s/domain/entity"
	"kill8s/infrastructure/persistence"
)

var _ Repository = new(RepoRedis)

type RepoRedis struct {
	DB *redis.Client
}

func NewRepoRedis() *RepoRedis {
	return &RepoRedis{
		DB: persistence.GetRedis(),
	}
}

func (repo *RepoRedis) Register(ctx context.Context, user *entity.User) error {
	return nil
}

func (repo *RepoRedis) Login(ctx context.Context, user *entity.User) error {
	return nil
}
