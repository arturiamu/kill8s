package handler

import (
	"github.com/gin-gonic/gin"
)

// ApiNamespacedCreate
// /api/{api-version}/namespaces/{namespace-name}/{resource-kind}/{name}
// /api/v1/namespaces/{namespace}/pods/{name}
func (h *K8sHandler) ApiNamespacedCreate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApiNamespacedCreate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Creater(ctx, namespace, resource, name)
}

// ApiNoneNamespacedCreate
// /api/{api-version}/{resource-kind}/{name}
// /api/v1/nodes/{name}
func (h *K8sHandler) ApiNoneNamespacedCreate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApiNoneNamespacedCreate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Creater(ctx, namespace, resource, name)
}

// ApisNamespacedCreate
// /apis/{api-group}/{api-version}/namespaces/{namespace}/{resource-kind}/{name}
// /apps/v1/namespaces/{namespace-name}/daemonsets/{name}
func (h *K8sHandler) ApisNamespacedCreate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNamespacedCreate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Creater(ctx, namespace, resource, name)
}

// ApisNoneNamespacedCreate
// x
func (h *K8sHandler) ApisNoneNamespacedCreate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNamespacedCreate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Creater(ctx, namespace, resource, name)
}

// NoneNamespacedUrlCreate
// x
func (h *K8sHandler) NoneNamespacedUrlCreate(ctx *gin.Context) (res any, err error) {
	h.logger.Println("K8sHandler.ApisNamespacedCreate")
	namespace := h.getNamespaceName(ctx)
	resource := h.getResourceKind(ctx)
	name := h.getResourceName(ctx)
	return h.application.Creater(ctx, namespace, resource, name)
}
