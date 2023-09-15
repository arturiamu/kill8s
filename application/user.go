package application

import (
	"context"
	"github.com/pkg/errors"
	userRepo "kill8s/domain/repository/user"
	"kill8s/domain/service"
	userDto "kill8s/server/dto/user"
	"log"
)

type UserApplication struct {
	logger         *log.Logger
	userRepository userRepo.Repository
	userService    *service.UserService
}

func NewUserApplication(l *log.Logger) *UserApplication {
	ua := &UserApplication{
		logger:         l,
		userRepository: userRepo.GetRepository(),
		userService:    service.NewUserService(),
	}
	//ua.logger.SetPrefix("UserApplication:")
	return ua
}

func (app *UserApplication) Register(ctx context.Context, req *userDto.RegisterReq) (*userDto.RegisterResp, error) {
	// application中通过调用各个service, 实现业务逻辑的编排；
	// 若业务逻辑过于简单（如简单的增删查改），可越过domain层，在application中直接使用repository进行操作
	user, err := app.userService.DoRegister(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "test post err")
	}

	return &userDto.RegisterResp{
		User: user,
	}, nil
}

func (app *UserApplication) Login(ctx context.Context, req *userDto.RegisterReq) (*userDto.RegisterResp, error) {
	user, err := app.userService.DoLogin(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "test post err")
	}

	return &userDto.RegisterResp{
		User: user,
	}, nil
}
