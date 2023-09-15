package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// ApiNamespacedList
// /api/{api-version}/namespaces/{namespace-name}/{resource-kind}
// /api/v1/namespaces/{namespace}/pods
func (h *K8sHandler) ApiNamespacedList(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApiNamespacedList")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	if namespace == "" || resource == "" || name != "" {
		return nil, errors.New("the request parameter is incorrect")
	}
	return h.application.Lister(ctx, namespace, resource)
}

// ApiNoneNamespacedList
// /api/{api-version}/{resource-kind}
// /api/v1/namespaces
// /api/v1/nodes
func (h *K8sHandler) ApiNoneNamespacedList(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApiNoneNamespacedList")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	if namespace != "" || resource == "" || name != "" {
		return nil, errors.New("the request parameter is incorrect")
	}
	return h.application.Lister(ctx, namespace, resource)
}

// ApisNamespacedList
// /apis/{api-group}/{api-version}/namespaces/{namespace-name}/{resource-kind}
// /apis/apps/v1/namespaces/{namespace}/daemonsets
func (h *K8sHandler) ApisNamespacedList(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNamespacedList")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	if namespace == "" || resource == "" || name != "" {
		return nil, errors.New("the request parameter is incorrect")
	}
	return h.application.Lister(ctx, namespace, resource)
}

// ApisNoneNamespacedList
// /apis/{api-group}/{api-version}/{resource-kind}
// /apis/apps/v1/replicasets
func (h *K8sHandler) ApisNoneNamespacedList(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNoneNamespacedList")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	if namespace != "" || resource == "" || name != "" {
		return nil, errors.New("the request parameter is incorrect")
	}
	return h.application.Lister(ctx, namespace, resource)
}

func (h *K8sHandler) NoneNamespacedUrlList(ctx *gin.Context) (res any, err error) {
	return
}
