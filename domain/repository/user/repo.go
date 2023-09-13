package user

import (
	"kill8s/domain/entity"
)

var globalRepository Repository

type Repository interface {
	AddUser(user entity.User) (u entity.User, err error)
	DeleteUser(user entity.User) (u entity.User, err error)
	UpdateUser(user entity.User) (u entity.User, err error)
	GetUserByID(id int64) (user entity.User, err error)
	GetUserByName(name string) (user entity.User, err error)
	ListUser() (users []entity.User, err error)
}

func GetRepository() Repository {
	return globalRepository
}

func SetRepository(repo Repository) {
	globalRepository = repo
}
