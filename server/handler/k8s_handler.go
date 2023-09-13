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

// Init router("/api/v1")
func (h *K8sHandler) Init(router *gin.RouterGroup) {
	h.logger = log.Default()
	h.application = application.NewK8sApplication(h.logger)
	h.logger.SetPrefix("K8sHandler: ")

	if router != nil {
		k8sRouter := router.Group("/k8s", middleware.K8s())
		apiRouter := k8sRouter.Group("/api")
		h.ApiRouter(apiRouter)
		apisRouter := k8sRouter.Group("/apis")
		h.ApisRouter(apisRouter)
		h.InitNaneNamespacedUrl(k8sRouter)
	}
}

// ApiRouter /api
// /v1/namespaces/{namespace}/pods
// /v1/namespaces/{namespace}/pods/{{name}}
// /v1/namespaces/{namespace}/events
// /v1/namespaces/{namespace}/events
// /v1/events
// /v1/nodes
// /v1/nodes/{name}
func (h *K8sHandler) ApiRouter(router *gin.RouterGroup) {
	namespacedRouter := router.Group("/:api-version/namespaces")
	{
		// /api/v1/namespaces/{namespace}/pods
		namespacedRouter.GET("/:namespace-name/:resource-kind")
		// /api/v1/namespaces/{namespace}/pods/{{name}}
		namespacedRouter.GET("/:namespace-name/:resource-kind/:resource-name")
	}

	naneNamespacedRouter := router.Group("/:api-version")
	{
		// /api/v1/nodes
		naneNamespacedRouter.GET("/:resource-kind")
		// /api/v1/nodes/{name}
		naneNamespacedRouter.GET("/:resource-kind/:resource-name")
	}
}

// ApisRouter /apis
// /apps/v1/namespaces/{namespace}/daemonsets
// /apps/v1/namespaces/{namespace}/daemonsets/{{name}}
// /batch/v1/namespaces/{namespace}/jobs
// /batch/v1/namespaces/{namespace}/jobs/{name}
// /batch/v1/jobs
// /apps/v1/replicasets
func (h *K8sHandler) ApisRouter(router *gin.RouterGroup) {
	namespacedRouter := router.Group("/:api-group/:api-version/namespaces")
	{
		// /apis/apps/v1/namespaces/{namespace}/daemonsets
		namespacedRouter.GET("/:namespace-name/:resource-kind")
		// /apps/v1/namespaces/{namespace}/daemonsets/{{name}}
		namespacedRouter.GET("/:namespace-name/:resource-kind/:resource-name")
	}

	naneNamespacedRouter := router.Group("/:api-group/:api-version")
	{
		naneNamespacedRouter.GET("/:resource-kind")
		naneNamespacedRouter.GET("/:resource-kind/:resource-name")
	}
}

// InitNaneNamespacedUrl http://localhost:8080/apis/rbac.authorization.k8s.io/v1/clusterroles/test-clusterrole
func (h *K8sHandler) InitNaneNamespacedUrl(router *gin.RouterGroup) {
	router.GET("/healthz", middleware.Wrapper(h.Todo))
	router.GET("/healthz/:type", middleware.Wrapper(h.Todo))
	router.GET("/healthz/poststarthook/:type", middleware.Wrapper(h.Todo))
	router.GET("/logs", middleware.Wrapper(h.Todo))
	router.GET("/metrics", middleware.Wrapper(h.Todo))
	router.GET("/openapi/v2", middleware.Wrapper(h.Todo))
	router.GET("/swagger-2.0.0.json", middleware.Wrapper(h.Todo))
	router.GET("/swagger-2.0.0.pb-v1", middleware.Wrapper(h.Todo))
	router.GET("/swagger-2.0.0.pb-v1.gz", middleware.Wrapper(h.Todo))
	router.GET("/swagger.json", middleware.Wrapper(h.Todo))
	router.GET("/swaggerapi", middleware.Wrapper(h.Todo))
	router.GET("/version", middleware.Wrapper(h.Todo))
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

func (h *K8sHandler) Todo(ctx *gin.Context) (resp any, err error) {
	return "not yet implemented", nil
}
