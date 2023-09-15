package handler

import (
	"github.com/gin-gonic/gin"
	"kill8s/application"
	k8sRepo "kill8s/domain/repository/k8s"
	"kill8s/infrastructure/constant"
	"kill8s/server/middleware"
	"log"
)

type K8sHandler struct {
	logger      *log.Logger
	application k8sRepo.Repository
}

// Init router("/api/v1")
func (h *K8sHandler) Init(router *gin.RouterGroup) {
	h.logger = log.Default()
	//h.application = application.NewK8sApplication(h.logger)
	h.application = application.NewK8sApplicationMock(h.logger)
	//h.logger.SetPrefix("K8sHandler: ")

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
// /v1/events
// /v1/nodes
// /v1/nodes/{name}
func (h *K8sHandler) ApiRouter(router *gin.RouterGroup) {
	namespacedRouter := router.Group("/:api-version/namespaces")
	{
		// /api/v1/namespaces/{namespace}/pods
		namespacedRouter.POST("/:namespace-name/:resource-kind", middleware.Wrapper(h.ApiNamespacedCreate))

		// /api/v1/namespaces/{namespace}/pods/{name}
		namespacedRouter.DELETE("/:namespace-name/:resource-kind/:resource-name", middleware.Wrapper(h.ApiNamespacedDelete))
		namespacedRouter.PATCH("/:namespace-name/:resource-kind/:resource-name", middleware.Wrapper(h.ApiNamespacedUpdate))
		namespacedRouter.GET("/:namespace-name/:resource-kind/:resource-name", middleware.Wrapper(h.ApiNamespacedGet))

		// /api/v1/namespaces/{namespace}/pods
		namespacedRouter.GET("/:namespace-name/:resource-kind", middleware.Wrapper(h.ApiNamespacedList))
	}

	naneNamespacedRouter := router.Group("/:api-version")
	{
		// /api/v1/namespaces
		// /api/v1/nodes
		naneNamespacedRouter.POST("/:resource-kind", middleware.Wrapper(h.ApiNoneNamespacedCreate))

		// /api/v1/namespaces/{name}
		// /api/v1/nodes/{name}
		naneNamespacedRouter.DELETE("/:resource-kind/:resource-name", middleware.Wrapper(h.ApiNoneNamespacedDelete))
		naneNamespacedRouter.PATCH("/:resource-kind/:resource-name", middleware.Wrapper(h.ApiNoneNamespacedUpdate))
		naneNamespacedRouter.GET("/:resource-kind/:resource-name", middleware.Wrapper(h.ApiNoneNamespacedGet))

		// /api/v1/nodes
		// /api/v1/namespaces
		naneNamespacedRouter.GET("/:resource-kind", middleware.Wrapper(h.ApiNoneNamespacedList))
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
		// /apis/batch/v1/namespaces/{namespace}/jobs
		namespacedRouter.POST("/:namespace-name/:resource-kind", middleware.Wrapper(h.ApisNamespacedCreate))

		// /apis/batch/v1/namespaces/{namespace}/jobs/{name}
		namespacedRouter.DELETE("/:namespace-name/:resource-kind/:resource-name", middleware.Wrapper(h.ApisNamespacedDelete))
		namespacedRouter.PATCH("/:namespace-name/:resource-kind/:resource-name", middleware.Wrapper(h.ApisNamespacedUpdate))
		namespacedRouter.GET("/:namespace-name/:resource-kind/:resource-name", middleware.Wrapper(h.ApisNamespacedGet))

		// /apis/batch/v1/namespaces/{namespace}/jobs
		namespacedRouter.GET("/:namespace-name/:resource-kind", middleware.Wrapper(h.ApisNamespacedList))
	}

	naneNamespacedRouter := router.Group("/:api-group/:api-version")
	{
		//naneNamespacedRouter.POST("/:resource-kind", middleware.Wrapper(h.ApisNoneNamespacedCreate))

		// /apis/batch/v1/jobs
		//naneNamespacedRouter.DELETE("/:resource-kind/:resource-name", middleware.Wrapper(h.ApisNoneNamespacedDelete))
		//naneNamespacedRouter.PATCH("/:resource-kind/:resource-name", middleware.Wrapper(h.ApisNoneNamespacedUpdate))
		//naneNamespacedRouter.GET("/:resource-kind/:resource-name", middleware.Wrapper(h.ApisNoneNamespacedList))
		// /batch/v1/jobs
		naneNamespacedRouter.GET("/:resource-kind", middleware.Wrapper(h.ApisNoneNamespacedList))
	}
}

// InitNaneNamespacedUrl http://localhost:8080/healthz
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
