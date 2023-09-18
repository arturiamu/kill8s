package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kill8s/infrastructure/constant"
	"kill8s/server/dto"
	"net/http"
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

		resource := constant.NewSupportedResources(resourceKind)
		if !resource.ResourceCheck() {
			context.JSON(http.StatusInternalServerError, dto.GetErrorResponse("unsupported resource type"))
			_ = context.AbortWithError(http.StatusInternalServerError, errors.New("unsupported resource type"))
			return
		}
		var act constant.SupportedResourcesAction
		switch context.Request.Method {
		case http.MethodGet:
			act = constant.ActionGet
		case http.MethodPost:
			act = constant.ActionCreate
		case http.MethodPatch:
			act = constant.ActionUpdate
		case http.MethodPut:
			act = constant.ActionUpdate
		case http.MethodDelete:
			act = constant.ActionDelete
		}
		if !resource.ActionCheck(act) {
			context.JSON(http.StatusInternalServerError, dto.GetErrorResponse("unsupported resource action"))
			_ = context.AbortWithError(http.StatusInternalServerError, errors.New("unsupported resource action"))
			return
		}
		context.Next()
	}
}
