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
	ua.logger.SetPrefix("UserApplication:")
	return ua
}

func (app *UserApplication) UserRegister(ctx context.Context, req *userDto.RegisterReq) (*userDto.RegisterResp, error) {
	user, err := app.userService.UserRegister(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "test post err")
	}

	return &userDto.RegisterResp{
		User: user,
	}, nil
}
