package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kill8s/infrastructure/kube"
)

var _ Repository = new(RepoClientGo)
var _ CoreV1 = new(RepoClientGo)

var _ AppsV1 = new(RepoClientGo)

type RepoClientGo struct {
	CS *kubernetes.Clientset
}

func (r *RepoClientGo) CoreV1() CoreV1 {
	return NewRepoClientGo()
}

func (r *RepoClientGo) AppsV1() AppsV1 {
	return NewRepoClientGo()
}

func NewRepoClientGo() *RepoClientGo {
	return &RepoClientGo{
		CS: kube.GetK8sClientSet(),
	}
}

func (r *RepoClientGo) DeleteControllerRevision(ctx context.Context, namespace, name string) error {
	return r.CS.AppsV1().ControllerRevisions(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetControllerRevision(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.AppsV1().ControllerRevisions(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListControllerRevision(ctx context.Context, namespace string) (any, error) {
	return r.CS.AppsV1().ControllerRevisions(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeleteDaemonSet(ctx context.Context, namespace, name string) error {
	return r.CS.AppsV1().DaemonSets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetDaemonSet(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.AppsV1().DaemonSets(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListDaemonSet(ctx context.Context, namespace string) (any, error) {
	return r.CS.AppsV1().DaemonSets(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeleteDeployment(ctx context.Context, namespace, name string) error {
	return r.CS.AppsV1().Deployments(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetDeployment(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListDeployment(ctx context.Context, namespace string) (any, error) {
	return r.CS.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeleteReplicaSet(ctx context.Context, namespace, name string) error {
	return r.CS.AppsV1().ReplicaSets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetReplicaSet(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.AppsV1().ReplicaSets(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListReplicaSet(ctx context.Context, namespace string) (any, error) {
	return r.CS.AppsV1().ReplicaSets(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeleteStatefulSet(ctx context.Context, namespace, name string) error {
	return r.CS.AppsV1().StatefulSets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetStatefulSet(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.AppsV1().StatefulSets(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListStatefulSet(ctx context.Context, namespace string) (any, error) {
	return r.CS.AppsV1().StatefulSets(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) CreateNamespace(ctx context.Context, name string) (any, error) {
	return r.CS.CoreV1().Namespaces().Create(ctx, &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}, metav1.CreateOptions{})
}

func (r *RepoClientGo) DeleteNamespace(ctx context.Context, name string) error {
	return r.CS.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetNamespace(ctx context.Context, name string) (any, error) {
	return r.CS.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListNamespace(ctx context.Context) (any, error) {
	return r.CS.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeleteService(ctx context.Context, namespace, name string) error {
	return r.CS.CoreV1().Services(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetService(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListService(ctx context.Context, namespace string) (any, error) {
	return r.CS.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeletePod(ctx context.Context, namespace, name string) error {
	return r.CS.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetPod(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListPod(ctx context.Context, namespace string) (any, error) {
	return r.CS.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeleteSecret(ctx context.Context, namespace, name string) error {
	return r.CS.CoreV1().Secrets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetSecret(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.CoreV1().Secrets(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListSecret(ctx context.Context, namespace string) (any, error) {
	return r.CS.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeletePersistentVolume(ctx context.Context, name string) error {
	return r.CS.CoreV1().PersistentVolumes().Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetPersistentVolume(ctx context.Context, name string) (any, error) {
	return r.CS.CoreV1().PersistentVolumes().Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListPersistentVolume(ctx context.Context) (any, error) {
	return r.CS.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeleteReplicationController(ctx context.Context, namespace, name string) error {
	return r.CS.CoreV1().ReplicationControllers(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetReplicationController(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.CoreV1().ReplicationControllers(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListReplicationController(ctx context.Context, namespace string) (any, error) {
	return r.CS.CoreV1().ReplicationControllers(namespace).List(ctx, metav1.ListOptions{})
}

func (r *RepoClientGo) DeleteConfigMap(ctx context.Context, namespace, name string) error {
	return r.CS.CoreV1().ConfigMaps(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (r *RepoClientGo) GetConfigMap(ctx context.Context, namespace, name string) (any, error) {
	return r.CS.CoreV1().ConfigMaps(namespace).Get(ctx, name, metav1.GetOptions{})
}

func (r *RepoClientGo) ListConfigMap(ctx context.Context, namespace string) (any, error) {
	return r.CS.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
}
