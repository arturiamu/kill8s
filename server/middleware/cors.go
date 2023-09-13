package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		//context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
		context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
		context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.Next()
	}
}
