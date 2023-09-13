package user

import (
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

func (repo *RepoMysql) AddUser(user entity.User) (u entity.User, err error) {
	err = repo.DB.Create(&user).Error
	return user, err
}

func (repo *RepoMysql) DeleteUser(user entity.User) (u entity.User, err error) {
	err = repo.DB.Delete(&user).Error
	return user, err
}

func (repo *RepoMysql) UpdateUser(user entity.User) (u entity.User, err error) {
	err = repo.DB.Updates(&user).Error
	return user, err
}

func (repo *RepoMysql) GetUserByID(id int64) (user entity.User, err error) {
	err = repo.DB.Find(&user).Error
	return
}

func (repo *RepoMysql) GetUserByName(name string) (user entity.User, err error) {
	err = repo.DB.Find(&user).Error
	return
}

func (repo *RepoMysql) ListUser() (users []entity.User, err error) {
	err = repo.DB.Find(&users).Error
	return
}
