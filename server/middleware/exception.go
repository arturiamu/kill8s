package middleware

import (
	"github.com/gin-gonic/gin"
	"kill8s/server/dto"
	"net/http"
)

type resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ExceptionHandlerFunc func(c *gin.Context) (data any, err error)

func Wrapper(handlerFunc ExceptionHandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := handlerFunc(c)
		if err != nil {
			c.JSON(http.StatusOK, dto.GetErrorResponse(err.Error()))
			return
		}
		c.JSON(http.StatusOK, dto.GetSuccessResponse(data))
	}
}
