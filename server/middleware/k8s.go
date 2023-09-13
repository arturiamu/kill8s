package middleware

import (
	"github.com/gin-gonic/gin"
	"kill8s/infrastructure/constant"
)

func K8s() gin.HandlerFunc {
	return func(context *gin.Context) {
		namespaceName := context.Param(constant.K8sNamespaceName)
		context.Set(constant.K8sNamespaceName, namespaceName)

		apiVersion := context.Param(constant.K8sApiVersion)
		context.Set(constant.K8sApiVersion, apiVersion)

		apiGroup := context.Param(constant.K8sApiGroup)
		context.Set(constant.K8sApiGroup, apiGroup)

		resourceKind := context.Param(constant.K8sResourceKind)
		context.Set(constant.K8sResourceKind, resourceKind)

		resourceName := context.Param(constant.K8sResourceName)
		context.Set(constant.K8sResourceName, resourceName)

		context.Next()
	}
}
