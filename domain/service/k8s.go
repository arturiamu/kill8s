package service

import (
	"context"
	k8sRepo "kill8s/domain/repository/k8s"
)

type K8sService struct {
	repo k8sRepo.Repository
}

func NewK8sService() *K8sService {
	return &K8sService{
		repo: k8sRepo.GetRepository(),
	}
}

func (service *K8sService) ListNamespaces(ctx context.Context) (resp any, err error) {
	return service.repo.ListNamespaces(ctx)
}
