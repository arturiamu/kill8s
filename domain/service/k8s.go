package service

import (
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
