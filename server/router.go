package server

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"kill8s/server/handler"
	"kill8s/server/middleware"
	"net/http"
)

func NewRouter() (router *gin.Engine) {
	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.Use(middleware.Cors())

	pprof.Register(router)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	router.GET("/ping", middleware.Wrapper(Ping))
	v1 := router.Group("/api/v1")
	// init handlers
	initHandlers(v1)
	router.NoRoute(noRouter)
	return router
}

type Handler interface {
	Init(*gin.RouterGroup)
}

func Ping(ctx *gin.Context) (resp any, err error) {
	return "pong", nil
}

func initHandlers(router *gin.RouterGroup) {
	handlers := []Handler{
		&handler.UserHandler{},
		&handler.K8sHandler{},
	}
	for _, hdl := range handlers {
		hdl.Init(router)
	}
}

func noRouter(ctx *gin.Context) {
	fmt.Println(1)
	fmt.Println(ctx.FullPath())
}
