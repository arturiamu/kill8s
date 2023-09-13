package k8s

var globalRepository Repository

type CoreV1 interface {
	NamespacesGetter
	ServicesGetter
	PodsGetter
	SecretsGetter
	PersistentVolumesGetter
	ReplicationControllersGetter
	ConfigMapsGetter
}

type AppsV1 interface {
	ControllerRevisionsGetter
	DaemonSetsGetter
	DeploymentsGetter
	ReplicaSetsGetter
	StatefulSetsGetter
}

type Repository interface {
	CoreV1() CoreV1
	AppsV1() AppsV1
}

func GetRepository() Repository {
	return globalRepository
}

func SetRepository(repo Repository) {
	globalRepository = repo
}
