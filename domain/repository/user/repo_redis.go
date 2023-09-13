package user

import (
	"github.com/go-redis/redis/v8"
	"kill8s/domain/entity"
	"kill8s/infrastructure/persistence"
)

var _ Repository = new(RepoRedis)

type RepoRedis struct {
	DB *redis.Client
}

func (r RepoRedis) AddUser(user entity.User) (u entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepoRedis) DeleteUser(user entity.User) (u entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepoRedis) UpdateUser(user entity.User) (u entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepoRedis) GetUserByID(id int64) (user entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepoRedis) GetUserByName(name string) (user entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepoRedis) ListUser() (users []entity.User, err error) {
	//TODO implement me
	panic("implement me")
}

func NewRepoRedis() *RepoRedis {
	return &RepoRedis{
		DB: persistence.GetRedis(),
	}
}
