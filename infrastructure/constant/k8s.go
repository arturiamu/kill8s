package constant

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	K8sNamespaces    = "namespaces"
	K8sNamespaceName = "namespace-name"
	K8sApiVersion    = "api-version"
	K8sApiGroup      = "api-group"
	K8sResourceKind  = "resource-kind"
	K8sResourceName  = "resource-name"
)

type SupportedResourcesAction string

const (
	K8sSupportedResourcesActionCreate                    SupportedResourcesAction = "create"
	K8sSupportedResourcesActionUpdate                    SupportedResourcesAction = "update"
	K8sSupportedResourcesActionUpdateStatus              SupportedResourcesAction = "updatestatus"
	K8sSupportedResourcesActionDelete                    SupportedResourcesAction = "delete"
	K8sSupportedResourcesActionDeleteCollection          SupportedResourcesAction = "deletecollection"
	K8sSupportedResourcesActionGet                       SupportedResourcesAction = "get"
	K8sSupportedResourcesActionList                      SupportedResourcesAction = "list"
	K8sSupportedResourcesActionWatch                     SupportedResourcesAction = "watch"
	K8sSupportedResourcesActionPatch                     SupportedResourcesAction = "patch"
	K8sSupportedResourcesActionApply                     SupportedResourcesAction = "apply"
	K8sSupportedResourcesActionApplyStatus               SupportedResourcesAction = "applystatus"
	K8sSupportedResourcesActionUpdateEphemeralContainers SupportedResourcesAction = "updateephemeralcontainers"
	K8sSupportedResourcesActionExpansion                 SupportedResourcesAction = "expansion"
)

func NewSupportedResourcesAction(act string) SupportedResourcesAction {
	return SupportedResourcesAction(act)
}

type SupportedResources string

const (
	SupportedResourcesPods                   SupportedResources = "pods"
	SupportedResourcesNodes                  SupportedResources = "nodes"
	SupportedResourcesSecrets                SupportedResources = "secrets"
	SupportedResourcesNamespaces             SupportedResources = "namespaces"
	SupportedResourcesServices               SupportedResources = "services"
	SupportedResourcesPersistentVolumes      SupportedResources = "persistentVolumes"
	SupportedResourcesReplicationControllers SupportedResources = "replicationcontrollers"
	SupportedResourcesConfigMaps             SupportedResources = "configmaps"

	SupportedResourcesControllerRevisions SupportedResources = "controllerrevisions"
	SupportedResourcesDaemonSets          SupportedResources = "daemonsets"
	SupportedResourcesDeployments         SupportedResources = "deployments"
	SupportedResourcesReplicaSets         SupportedResources = "replicasets"
	SupportedResourcesStatefulSets        SupportedResources = "statefulsets"
)

var supportedResourcesList = []SupportedResources{
	SupportedResourcesPods, SupportedResourcesSecrets, SupportedResourcesNamespaces, SupportedResourcesServices,
	SupportedResourcesPersistentVolumes, SupportedResourcesReplicationControllers, SupportedResourcesConfigMaps,
	SupportedResourcesControllerRevisions, SupportedResourcesDaemonSets, SupportedResourcesDeployments,
	SupportedResourcesReplicaSets, SupportedResourcesStatefulSets,
}

var supportedResourcesActionMap = map[SupportedResources]map[SupportedResourcesAction]struct{}{
	SupportedResourcesPods: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesSecrets: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesNamespaces: {
		K8sSupportedResourcesActionCreate: struct{}{},
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesServices: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesPersistentVolumes: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesReplicationControllers: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesConfigMaps: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesControllerRevisions: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesDaemonSets: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesDeployments: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesReplicaSets: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesStatefulSets: {
		K8sSupportedResourcesActionDelete: struct{}{},
		K8sSupportedResourcesActionGet:    struct{}{},
		K8sSupportedResourcesActionList:   struct{}{},
	},
	SupportedResourcesNodes: {
		K8sSupportedResourcesActionGet:  struct{}{},
		K8sSupportedResourcesActionList: struct{}{},
	},
}

func NewSupportedResources(str string) SupportedResources {
	return SupportedResources(str)
}

func ListSupportedResources() []SupportedResources {
	return supportedResourcesList
}

func (s SupportedResources) ResourceCheck() bool {
	_, ok := supportedResourcesActionMap[s]
	return ok
}

func (s SupportedResources) ActionCheck(act SupportedResourcesAction) bool {
	actions, ok := supportedResourcesActionMap[s]
	if !ok {
		return false
	}
	_, okk := actions[act]
	return okk
}

func GetRuntimeObj(resource string) (obj runtime.Object) {
	res := NewSupportedResources(resource)
	switch res {
	case SupportedResourcesPods:
		return &corev1.Pod{}
	case SupportedResourcesSecrets:
		return &corev1.Secret{}
	case SupportedResourcesNamespaces:
		return &corev1.Namespace{}
	case SupportedResourcesServices:
		return &corev1.Service{}
	case SupportedResourcesPersistentVolumes:
		return &corev1.PersistentVolume{}
	case SupportedResourcesReplicationControllers:
		return &corev1.ReplicationController{}
	case SupportedResourcesConfigMaps:
		return &corev1.ConfigMap{}
	case SupportedResourcesNodes:
		return &corev1.Node{}

	case SupportedResourcesControllerRevisions:
		return &appsv1.ControllerRevision{}
	case SupportedResourcesDaemonSets:
		return &appsv1.DaemonSet{}
	case SupportedResourcesDeployments:
		return &appsv1.Deployment{}
	case SupportedResourcesReplicaSets:
		return &appsv1.ReplicaSet{}
	case SupportedResourcesStatefulSets:
		return &appsv1.StatefulSet{}
	default:
		return nil
	}
}

func ListRuntimeObj(resource string) (obj runtime.Object) {
	res := NewSupportedResources(resource)
	switch res {
	case SupportedResourcesPods:
		return &corev1.PodList{}
	case SupportedResourcesSecrets:
		return &corev1.SecretList{}
	case SupportedResourcesNamespaces:
		return &corev1.NamespaceList{}
	case SupportedResourcesServices:
		return &corev1.ServiceList{}
	case SupportedResourcesPersistentVolumes:
		return &corev1.PersistentVolumeList{}
	case SupportedResourcesReplicationControllers:
		return &corev1.ReplicationControllerList{}
	case SupportedResourcesConfigMaps:
		return &corev1.ConfigMapList{}
	case SupportedResourcesNodes:
		return &corev1.NodeList{}

	case SupportedResourcesControllerRevisions:
		return &appsv1.ControllerRevisionList{}
	case SupportedResourcesDaemonSets:
		return &appsv1.DaemonSetList{}
	case SupportedResourcesDeployments:
		return &appsv1.DeploymentList{}
	case SupportedResourcesReplicaSets:
		return &appsv1.ReplicaSetList{}
	case SupportedResourcesStatefulSets:
		return &appsv1.StatefulSetList{}
	default:
		return nil
	}
}
