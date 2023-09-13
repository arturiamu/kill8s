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

func (service *UserService) UserRegister(ctx context.Context, name string, password string) (*entity.User, error) {
	user := &entity.User{
		Username: name,
		Password: password,
	}
	err := user.MD5Password()
	if err != nil {
		return nil, errors.Wrap(err, "md5 password err")
	}

	err = service.repo.Register(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "test repo register err")
	}

	return user, nil
}

func (service *UserService) UserLogin(ctx context.Context, name string, password string) (*entity.User, error) {
	user := &entity.User{
		Username: name,
		Password: password,
	}
	err := user.MD5Password()
	if err != nil {
		return nil, errors.Wrap(err, "md5 password err")
	}

	err = service.repo.Register(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "test repo register err")
	}

	return user, nil
}
