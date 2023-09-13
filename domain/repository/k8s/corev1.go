package k8s

import "context"

type NamespacesGetter interface {
	CreateNamespace(ctx context.Context, name string) (any, error)
	DeleteNamespace(ctx context.Context, name string) error
	GetNamespace(ctx context.Context, name string) (any, error)
	ListNamespace(ctx context.Context) (any, error)
}

type ServicesGetter interface {
	DeleteService(ctx context.Context, namespace, name string) error
	GetService(ctx context.Context, namespace, name string) (any, error)
	ListService(ctx context.Context, namespace string) (any, error)
}

type PodsGetter interface {
	DeletePod(ctx context.Context, namespace, name string) error
	GetPod(ctx context.Context, namespace, name string) (any, error)
	ListPod(ctx context.Context, namespace string) (any, error)
}

type SecretsGetter interface {
	DeleteSecret(ctx context.Context, namespace, name string) error
	GetSecret(ctx context.Context, namespace, name string) (any, error)
	ListSecret(ctx context.Context, namespace string) (any, error)
}

type PersistentVolumesGetter interface {
	DeletePersistentVolume(ctx context.Context, name string) error
	GetPersistentVolume(ctx context.Context, name string) (any, error)
	ListPersistentVolume(ctx context.Context) (any, error)
}

type ReplicationControllersGetter interface {
	DeleteReplicationController(ctx context.Context, namespace, name string) error
	GetReplicationController(ctx context.Context, namespace, name string) (any, error)
	ListReplicationController(ctx context.Context, namespace string) (any, error)
}

type ConfigMapsGetter interface {
	DeleteConfigMap(ctx context.Context, namespace, name string) error
	GetConfigMap(ctx context.Context, namespace, name string) (any, error)
	ListConfigMap(ctx context.Context, namespace string) (any, error)
}
