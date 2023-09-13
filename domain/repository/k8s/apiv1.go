package k8s

import "context"

type ControllerRevisionsGetter interface {
	DeleteControllerRevision(ctx context.Context, namespace, name string) error
	GetControllerRevision(ctx context.Context, namespace, name string) (any, error)
	ListControllerRevision(ctx context.Context, namespace string) (any, error)
}

type DaemonSetsGetter interface {
	DeleteDaemonSet(ctx context.Context, namespace, name string) error
	GetDaemonSet(ctx context.Context, namespace, name string) (any, error)
	ListDaemonSet(ctx context.Context, namespace string) (any, error)
}

type DeploymentsGetter interface {
	DeleteDeployment(ctx context.Context, namespace, name string) error
	GetDeployment(ctx context.Context, namespace, name string) (any, error)
	ListDeployment(ctx context.Context, namespace string) (any, error)
}

type ReplicaSetsGetter interface {
	DeleteReplicaSet(ctx context.Context, namespace, name string) error
	GetReplicaSet(ctx context.Context, namespace, name string) (any, error)
	ListReplicaSet(ctx context.Context, namespace string) (any, error)
}

type StatefulSetsGetter interface {
	DeleteStatefulSet(ctx context.Context, namespace, name string) error
	GetStatefulSet(ctx context.Context, namespace, name string) (any, error)
	ListStatefulSet(ctx context.Context, namespace string) (any, error)
}
