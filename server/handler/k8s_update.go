package handler

import (
	"github.com/gin-gonic/gin"
)

// ApiNamespacedUpdate
// /api/{api-version}/namespaces/{namespace-name}/{resource-kind}/{name}
// /api/v1/namespaces/{namespace}/pods/{name}
func (h *K8sHandler) ApiNamespacedUpdate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApiNamespacedUpdate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Updater(ctx, namespace, resource, name)
}

// ApiNoneNamespacedUpdate
// /api/{api-version}/{resource-kind}/{name}
// /api/v1/nodes/{name}
func (h *K8sHandler) ApiNoneNamespacedUpdate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApiNoneNamespacedUpdate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Updater(ctx, namespace, resource, name)
}

// ApisNamespacedUpdate
// /apis/{api-group}/{api-version}/namespaces/{namespace}/{resource-kind}/{name}
// /apps/v1/namespaces/{namespace-name}/daemonsets/{name}
func (h *K8sHandler) ApisNamespacedUpdate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNamespacedUpdate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Updater(ctx, namespace, resource, name)
}

// ApisNoneNamespacedUpdate
// x
func (h *K8sHandler) ApisNoneNamespacedUpdate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNoneNamespacedUpdate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Updater(ctx, namespace, resource, name)
}

// NoneNamespacedUrlUpdate
// x
func (h *K8sHandler) NoneNamespacedUrlUpdate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.NoneNamespacedUrlUpdate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Updater(ctx, namespace, resource, name)
}
