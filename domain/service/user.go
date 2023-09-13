package service

import (
	"context"
	"github.com/pkg/errors"
	"kill8s/domain/entity"
	userRepo "kill8s/domain/repository/user"
)

type UserService struct {
	repo userRepo.Repository
}

func NewUserService() *UserService {
	return &UserService{
		repo: userRepo.GetRepository(),
	}
}

func (service *UserService) DoRegister(ctx context.Context, name string, password string) (*entity.User, error) {
	// domain 中执行具体的业务逻辑，封装完整的数据，供application可直接使用；
	// 若业务逻辑过于简单（如简单的增删查改），可跳过domain层，在application中直接使用repository进行操作
	user := entity.User{
		Username: name,
		Password: password,
	}
	err := user.MD5Password()
	if err != nil {
		return nil, errors.Wrap(err, "md5 password err")
	}

	_, err = service.repo.AddUser(user)
	if err != nil {
		return nil, errors.Wrap(err, "test repo register err")
	}

	return &user, nil
}

func (service *UserService) DoLogin(ctx context.Context, name string, password string) (*entity.User, error) {
	user := &entity.User{
		Username: name,
		Password: password,
	}
	err := user.MD5Password()
	if err != nil {
		return nil, errors.Wrap(err, "md5 password err")
	}
	u, err := service.repo.GetUserByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "test repo register err")
	}
	if user.Password == password {
		return &u, nil
	}
	return nil, errors.New("err username or password")
}
