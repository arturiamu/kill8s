package k8s

import (
	"k8s.io/client-go/kubernetes"
	"kill8s/infrastructure/kube"
)

type RepoClientGo struct {
	CS *kubernetes.Clientset
}

func NewRepoClientGo() *RepoClientGo {
	return &RepoClientGo{
		CS: kube.GetK8sClientSet(),
	}
}
