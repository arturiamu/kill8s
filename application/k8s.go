package application

import (
	"context"
	k8sRepo "kill8s/domain/repository/k8s"
	"kill8s/domain/service"
	"log"
)

type K8sApplication struct {
	logger        *log.Logger
	k8sRepository k8sRepo.Repository
	k8sService    *service.K8sService
}

var _ k8sRepo.AppsV1 = new(K8sApplication)
var _ k8sRepo.CoreV1 = new(K8sApplication)

func NewK8sApplication(l *log.Logger) *K8sApplication {
	ka := &K8sApplication{
		logger:        l,
		k8sRepository: k8sRepo.GetRepository(),
		k8sService:    service.NewK8sService(),
	}
	ka.logger.SetPrefix("K8sApplication:")

	return ka
}

func (k K8sApplication) CreateNamespace(ctx context.Context, name string) (any, error) {
	return k.k8sRepository.CoreV1().GetNamespace(ctx, name)
}

func (k K8sApplication) DeleteNamespace(ctx context.Context, name string) error {
	return k.k8sRepository.CoreV1().DeleteNamespace(ctx, name)
}

func (k K8sApplication) GetNamespace(ctx context.Context, name string) (any, error) {
	return k.k8sRepository.CoreV1().GetNamespace(ctx, name)
}

func (k K8sApplication) ListNamespace(ctx context.Context) (any, error) {
	return k.k8sRepository.CoreV1().ListNamespace(ctx)
}

func (k K8sApplication) DeleteService(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.CoreV1().DeleteService(ctx, namespace, name)
}

func (k K8sApplication) GetService(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.CoreV1().GetService(ctx, namespace, name)
}

func (k K8sApplication) ListService(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.CoreV1().ListService(ctx, namespace)
}

func (k K8sApplication) DeletePod(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.CoreV1().DeletePod(ctx, namespace, name)
}

func (k K8sApplication) GetPod(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.CoreV1().GetPod(ctx, namespace, name)
}

func (k K8sApplication) ListPod(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.CoreV1().ListPod(ctx, namespace)
}

func (k K8sApplication) DeleteSecret(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.CoreV1().DeleteSecret(ctx, namespace, name)
}

func (k K8sApplication) GetSecret(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.CoreV1().GetSecret(ctx, namespace, name)
}

func (k K8sApplication) ListSecret(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.CoreV1().ListSecret(ctx, namespace)
}

func (k K8sApplication) DeletePersistentVolume(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.CoreV1().DeletePersistentVolume(ctx, namespace, name)
}

func (k K8sApplication) GetPersistentVolume(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.CoreV1().GetPersistentVolume(ctx, namespace, name)
}

func (k K8sApplication) ListPersistentVolume(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.CoreV1().ListPersistentVolume(ctx, namespace)
}

func (k K8sApplication) DeleteReplicationController(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.CoreV1().DeleteReplicationController(ctx, namespace, name)
}

func (k K8sApplication) GetReplicationController(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.CoreV1().GetReplicationController(ctx, namespace, name)
}

func (k K8sApplication) ListReplicationController(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.CoreV1().ListReplicationController(ctx, namespace)
}

func (k K8sApplication) DeleteConfigMap(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.CoreV1().DeleteConfigMap(ctx, namespace, name)
}

func (k K8sApplication) GetConfigMap(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.CoreV1().GetConfigMap(ctx, namespace, name)
}

func (k K8sApplication) ListConfigMap(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.CoreV1().ListConfigMap(ctx, namespace)
}

func (k K8sApplication) DeleteControllerRevision(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.AppsV1().DeleteControllerRevision(ctx, namespace, name)
}

func (k K8sApplication) GetControllerRevision(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.AppsV1().GetControllerRevision(ctx, namespace, name)
}

func (k K8sApplication) ListControllerRevision(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.AppsV1().ListControllerRevision(ctx, namespace)
}

func (k K8sApplication) DeleteDaemonSet(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.AppsV1().DeleteDaemonSet(ctx, namespace, name)
}

func (k K8sApplication) GetDaemonSet(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.AppsV1().GetDaemonSet(ctx, namespace, name)
}

func (k K8sApplication) ListDaemonSet(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.AppsV1().ListDaemonSet(ctx, namespace)
}

func (k K8sApplication) DeleteDeployment(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.AppsV1().DeleteDeployment(ctx, namespace, name)
}

func (k K8sApplication) GetDeployment(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.AppsV1().GetDeployment(ctx, namespace, name)
}

func (k K8sApplication) ListDeployment(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.AppsV1().ListDeployment(ctx, namespace)
}

func (k K8sApplication) DeleteReplicaSet(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.AppsV1().DeleteControllerRevision(ctx, namespace, name)
}

func (k K8sApplication) GetReplicaSet(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.AppsV1().GetReplicaSet(ctx, namespace, name)
}

func (k K8sApplication) ListReplicaSet(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.AppsV1().ListReplicaSet(ctx, namespace)
}

func (k K8sApplication) DeleteStatefulSet(ctx context.Context, namespace, name string) error {
	return k.k8sRepository.AppsV1().DeleteStatefulSet(ctx, namespace, name)
}

func (k K8sApplication) GetStatefulSet(ctx context.Context, namespace, name string) (any, error) {
	return k.k8sRepository.AppsV1().GetStatefulSet(ctx, namespace, name)
}

func (k K8sApplication) ListStatefulSet(ctx context.Context, namespace string) (any, error) {
	return k.k8sRepository.AppsV1().ListStatefulSet(ctx, namespace)
}
