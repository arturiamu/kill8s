package handler

import (
	"github.com/gin-gonic/gin"
	"kill8s/application"
	"kill8s/infrastructure/constant"
	"kill8s/server/middleware"
	"log"
)

type K8sHandler struct {
	logger      *log.Logger
	application *application.K8sApplication
}

func (h *K8sHandler) Init(router *gin.RouterGroup) {
	h.logger = log.Default()
	h.application = application.NewK8sApplication(h.logger)
	h.logger.SetPrefix("K8sHandler: ")

	if router != nil {
		k8sRouter := router.Group("/k8s", middleware.K8s())
		namespaced := k8sRouter.Group("/api")
		h.InitNamespaced(namespaced)
		noneNamespaced := k8sRouter.Group("/apis")
		h.InitNaneNamespaced(noneNamespaced)
		h.InitNaneNamespacedUrl(k8sRouter)
	}
}

// InitNamespaced http://localhost:8080/api /v1/namespaces/default/pods/test-pod
func (h *K8sHandler) InitNamespaced(router *gin.RouterGroup) {
	router.GET("/:api-version/namespaces/:namespace-name/:resource-kind", middleware.Wrapper(h.NamespacedList))
	router.GET("/:api-version/namespaces/:namespace-name/:resource-kind/:resource-name",
		middleware.Wrapper(h.NamespacedGet))
}

// InitNaneNamespaced http://localhost:8080/apis /rbac.authorization.k8s.io/v1/clusterroles/test-clusterrole
func (h *K8sHandler) InitNaneNamespaced(router *gin.RouterGroup) {
	router.GET("/:api-group/:api-version/:resource-kind", middleware.Wrapper(h.NoneNamespacedList))
	router.GET("/:api-group/:api-version/:resource-kind/:resource-name",
		middleware.Wrapper(h.NoneNamespacedGet))
}

// InitNaneNamespacedUrl http://localhost:8080/apis/rbac.authorization.k8s.io/v1/clusterroles/test-clusterrole
func (h *K8sHandler) InitNaneNamespacedUrl(router *gin.RouterGroup) {
	router.GET("/healthz")
	router.GET("/healthz/:type")
	router.GET("/healthz/poststarthook/:type")
	router.GET("/logs")
	router.GET("/metrics")
	router.GET("/openapi/v2")
	router.GET("/swagger-2.0.0.json")
	router.GET("/swagger-2.0.0.pb-v1")
	router.GET("/swagger-2.0.0.pb-v1.gz")
	router.GET("/swagger.json")
	router.GET("/swaggerapi")
	router.GET("/version")
}

func (h *K8sHandler) getNamespaceName(ctx *gin.Context) string {
	ns, _ := ctx.Get(constant.K8sNamespaceName)
	return ns.(string)
}

func (h *K8sHandler) getApiVersion(ctx *gin.Context) string {
	ns, _ := ctx.Get(constant.K8sApiVersion)
	return ns.(string)
}

func (h *K8sHandler) getApiGroup(ctx *gin.Context) string {
	ns, _ := ctx.Get(constant.K8sApiGroup)
	return ns.(string)
}

func (h *K8sHandler) getResourceKind(ctx *gin.Context) string {
	ns, _ := ctx.Get(constant.K8sResourceKind)
	return ns.(string)
}

func (h *K8sHandler) getResourceName(ctx *gin.Context) string {
	ns, _ := ctx.Get(constant.K8sResourceName)
	return ns.(string)
}
