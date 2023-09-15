package constant

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
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
	ActionCreate                    SupportedResourcesAction = "create"
	ActionUpdate                    SupportedResourcesAction = "update"
	ActionUpdateStatus              SupportedResourcesAction = "updatestatus"
	ActionDelete                    SupportedResourcesAction = "delete"
	ActionDeleteCollection          SupportedResourcesAction = "deletecollection"
	ActionGet                       SupportedResourcesAction = "get"
	ActionList                      SupportedResourcesAction = "list"
	ActionWatch                     SupportedResourcesAction = "watch"
	ActionPatch                     SupportedResourcesAction = "patch"
	ActionApply                     SupportedResourcesAction = "apply"
	ActionApplyStatus               SupportedResourcesAction = "applystatus"
	ActionUpdateEphemeralContainers SupportedResourcesAction = "updateephemeralcontainers"
	ActionExpansion                 SupportedResourcesAction = "expansion"
)

func NewAction(act string) SupportedResourcesAction {
	return SupportedResourcesAction(act)
}

type SupportedResources string

const (
	Nodes                  = "nodes"
	Namespaces             = "namespaces"
	ResourceQuotas         = "resourcequotas"
	Pods                   = "pods"
	ReplicationControllers = "replicationcontrollers"
	Endpoints              = "endpoints"
	Services               = "services"
	PersistentVolumes      = "persistentvolumes"
	PersistentVolumeClaims = "persistentvolumeclaims"
	ConfigMaps             = "configmaps"
	Secrets                = "secrets"
	Events                 = "events" /////////////// api-apis
	Deployments            = "deployments"
	StatefulSets           = "statefulsets"
	DaemonSets             = "daemonsets"
	Jobs                   = "jobs"
	CronJobs               = "cronjobs"
	ReplicaSets            = "replicasets"
	Ingress                = "ingress"
	NetworkPolicies        = "networkpolicies"
	ClusterRoles           = "clusterroles"
	ClusterRoleBindings    = "clusterrolebindings"
	Roles                  = "roles"
	RoleBindings           = "rolebindings"
)

var supportedResourcesList = []SupportedResources{
	Nodes, Namespaces, ResourceQuotas, Pods, ReplicationControllers, Endpoints, Services, PersistentVolumes,
	PersistentVolumeClaims, ConfigMaps, Secrets, Events,
	/////////////// api-apis
	Deployments, StatefulSets, DaemonSets, Jobs, CronJobs, ReplicaSets, Ingress, NetworkPolicies, ClusterRoles,
	ClusterRoleBindings, Roles, RoleBindings,
}

var ActionMap = map[SupportedResources]map[SupportedResourcesAction]struct{}{
	Pods: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	Secrets: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	Namespaces: {
		ActionCreate: struct{}{},
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	Services: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	PersistentVolumes: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	ReplicationControllers: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	ConfigMaps: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	DaemonSets: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	Deployments: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	ReplicaSets: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	StatefulSets: {
		ActionDelete: struct{}{},
		ActionGet:    struct{}{},
		ActionList:   struct{}{},
	},
	Nodes: {
		ActionGet:  struct{}{},
		ActionList: struct{}{},
	},
	ResourceQuotas: {
		ActionGet:  struct{}{},
		ActionList: struct{}{},
	},
}

func NewSupportedResources(str string) SupportedResources {
	return SupportedResources(str)
}

func ListSupportedResources() []SupportedResources {
	return supportedResourcesList
}

func (s SupportedResources) ResourceCheck() bool {
	_, ok := ActionMap[s]
	return ok
}

func (s SupportedResources) ActionCheck(act SupportedResourcesAction) bool {
	actions, ok := ActionMap[s]
	if !ok {
		return false
	}
	_, okk := actions[act]
	return okk
}

func GetRuntimeObj(resource string) (obj runtime.Object) {
	res := NewSupportedResources(resource)
	switch res {
	case Nodes:
		return &corev1.Node{}
	case Namespaces:
		return &corev1.Namespace{}
	case ResourceQuotas:
		return &corev1.ResourceQuota{}
	case Pods:
		return &corev1.Pod{}
	case ReplicationControllers:
		return &corev1.ReplicationController{}
	case Endpoints:
		return &corev1.Endpoints{}
	case Services:
		return &corev1.Service{}
	case PersistentVolumes:
		return &corev1.PersistentVolume{}
	case PersistentVolumeClaims:
		return &corev1.PersistentVolumeClaim{}
	case ConfigMaps:
		return &corev1.ConfigMap{}
	case Secrets:
		return &corev1.Secret{}
	case Events:
		return &corev1.Event{}
	case Deployments:
		return &appsv1.Deployment{}
	case StatefulSets:
		return &appsv1.StatefulSet{}
	case DaemonSets:
		return &appsv1.DaemonSet{}
	case Jobs:
		return &batchv1.Job{}
	case CronJobs:
		return &batchv1.CronJob{}
	case ReplicaSets:
		return &appsv1.ReplicaSet{}
	case Ingress:
		return &netv1.Ingress{}
	case NetworkPolicies:
		return &netv1.NetworkPolicy{}
	case ClusterRoles:
		return &rbacv1.ClusterRole{}
	case ClusterRoleBindings:
		return &rbacv1.ClusterRoleBinding{}
	case Roles:
		return &rbacv1.Role{}
	case RoleBindings:
		return &rbacv1.RoleBinding{}
	default:
		return nil
	}
}

func ListRuntimeObj(resource string) (obj runtime.Object) {
	res := NewSupportedResources(resource)
	switch res {
	case Nodes:
		return &corev1.NodeList{}
	case Namespaces:
		return &corev1.NamespaceList{}
	case ResourceQuotas:
		return &corev1.ResourceQuotaList{}
	case Pods:
		return &corev1.PodList{}
	case ReplicationControllers:
		return &corev1.ReplicationControllerList{}
	case Endpoints:
		return &corev1.EndpointsList{}
	case Services:
		return &corev1.ServiceList{}
	case PersistentVolumes:
		return &corev1.PersistentVolumeList{}
	case PersistentVolumeClaims:
		return &corev1.PersistentVolumeClaimList{}
	case ConfigMaps:
		return &corev1.ConfigMapList{}
	case Secrets:
		return &corev1.SecretList{}
	case Events:
		return &corev1.EventList{}
	case Deployments:
		return &appsv1.DeploymentList{}
	case StatefulSets:
		return &appsv1.StatefulSetList{}
	case DaemonSets:
		return &appsv1.DaemonSetList{}
	case Jobs:
		return &batchv1.JobList{}
	case CronJobs:
		return &batchv1.CronJobList{}
	case ReplicaSets:
		return &appsv1.ReplicaSetList{}
	case Ingress:
		return &netv1.IngressList{}
	case NetworkPolicies:
		return &netv1.NetworkPolicyList{}
	case ClusterRoles:
		return &rbacv1.ClusterRoleList{}
	case ClusterRoleBindings:
		return &rbacv1.ClusterRoleBindingList{}
	case Roles:
		return &rbacv1.RoleList{}
	case RoleBindings:
		return &rbacv1.RoleBindingList{}
	default:
		return nil
	}
}
