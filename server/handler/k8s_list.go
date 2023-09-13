package handler

import "github.com/gin-gonic/gin"

// ApiNamespacedList
// /api/{api-version}/namespaces/{namespace-name}/{resource-kind}
// /api/v1/namespaces/{namespace}/pods
func (h *K8sHandler) ApiNamespacedList(ctx *gin.Context) (res any, err error) {
	return
}

// ApiNoneNamespacedList
// /api/{api-version}/{resource-kind}
// /api/v1/namespaces
// /api/v1/nodes
func (h *K8sHandler) ApiNoneNamespacedList(ctx *gin.Context) (res any, err error) {
	kind := h.getResourceKind(ctx)
	if kind == "namespaces" {
		return h.application.ListNamespace(ctx)
	}
	return
}

// ApisNamespacedList
// /apis/{api-group}/{api-version}/namespaces/{namespace-name}/{resource-kind}
// /apis/apps/v1/namespaces/{namespace}/daemonsets
func (h *K8sHandler) ApisNamespacedList(ctx *gin.Context) (res any, err error) {
	return
}

// ApisNoneNamespacedList
// /apis/{api-group}/{api-version}/{resource-kind}
// /apis/apps/v1/replicasets
func (h *K8sHandler) ApisNoneNamespacedList(ctx *gin.Context) (res any, err error) {
	return
}

func (h *K8sHandler) NoneNamespacedUrlList(ctx *gin.Context) (res any, err error) {
	return
}
