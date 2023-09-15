package handler

import (
	"github.com/gin-gonic/gin"
)

// ApiNamespacedDelete
// /api/{api-version}/namespaces/{namespace-name}/{resource-kind}/{name}
// /api/v1/namespaces/{namespace}/pods/{name}
func (h *K8sHandler) ApiNamespacedDelete(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApiNamespacedDelete")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Deleter(ctx, namespace, resource, name)
}

// ApiNoneNamespacedDelete
// /api/{api-version}/{resource-kind}/{name}
// /api/v1/nodes/{name}
func (h *K8sHandler) ApiNoneNamespacedDelete(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApiNoneNamespacedDelete")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Deleter(ctx, namespace, resource, name)
}

// ApisNamespacedDelete
// /apis/{api-group}/{api-version}/namespaces/{namespace}/{resource-kind}/{name}
// /apps/v1/namespaces/{namespace-name}/daemonsets/{name}
func (h *K8sHandler) ApisNamespacedDelete(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNamespacedDelete")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Deleter(ctx, namespace, resource, name)
}

// ApisNoneNamespacedDelete
// x
func (h *K8sHandler) ApisNoneNamespacedDelete(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNoneNamespacedDelete")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Deleter(ctx, namespace, resource, name)
}

// NoneNamespacedUrlDelete
// x
func (h *K8sHandler) NoneNamespacedUrlDelete(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.NoneNamespacedUrlDelete")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Deleter(ctx, namespace, resource, name)
}
