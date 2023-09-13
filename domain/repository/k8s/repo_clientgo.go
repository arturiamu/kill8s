package k8s

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"kill8s/infrastructure/kube"
)

var _ Repository = new(RepoClientGo)

type RepoClientGo struct {
	CS *kubernetes.Clientset
}

func NewRepoClientGo() *RepoClientGo {
	return &RepoClientGo{
		CS: kube.GetK8sClientSet(),
	}
}

func (service *RepoClientGo) ListNamespaces(ctx context.Context) (resp any, err error) {
	return
}
