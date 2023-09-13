package handler

import (
	"github.com/gin-gonic/gin"
	"kill8s/application"
	"kill8s/server/dto/user"
	"kill8s/server/middleware"
	"log"
)

type UserHandler struct {
	logger      *log.Logger
	application *application.UserApplication
}

func (h *UserHandler) Init(router *gin.RouterGroup) {
	h.logger = log.Default()
	h.logger.SetPrefix("UserHandler: ")
	h.application = application.NewUserApplication(h.logger)
	if router != nil {
		userGroup := router.Group("/user")
		userGroup.POST("/register", middleware.Wrapper(h.RegisterHandler))
		userGroup.POST("/login", middleware.Wrapper(h.LoginHandler))
	}
}

func (h *UserHandler) RegisterHandler(ctx *gin.Context) (resp any, err error) {
	userDto := user.RegisterReq{}
	return h.application.Register(ctx, &userDto)
}

func (h *UserHandler) LoginHandler(ctx *gin.Context) (resp any, err error) {
	userDto := user.RegisterReq{}
	return h.application.Login(ctx, &userDto)
}
