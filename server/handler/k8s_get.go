package handler

import "github.com/gin-gonic/gin"

// ApiNamespacedGet
// /api/{api-version}/namespaces/{namespace-name}/{resource-kind}/{name}
// /api/v1/namespaces/{namespace}/pods/{name}
func (h *K8sHandler) ApiNamespacedGet(ctx *gin.Context) (res any, err error) {
	return
}

// ApiNoneNamespacedGet
// /api/{api-version}/{resource-kind}/{name}
// /api/v1/nodes/{name}
func (h *K8sHandler) ApiNoneNamespacedGet(ctx *gin.Context) (res any, err error) {
	return
}

// ApisNamespacedGet
// /apis/{api-group}/{api-version}/namespaces/{namespace}/{resource-kind}/{name}
// /apps/v1/namespaces/{namespace-name}/daemonsets/{name}
func (h *K8sHandler) ApisNamespacedGet(ctx *gin.Context) (res any, err error) {
	return
}

// ApisNoneNamespacedGet
// x
func (h *K8sHandler) ApisNoneNamespacedGet(ctx *gin.Context) (res any, err error) {
	return
}

// NoneNamespacedUrlGet
// x
func (h *K8sHandler) NoneNamespacedUrlGet(ctx *gin.Context) (res any, err error) {
	return
}
