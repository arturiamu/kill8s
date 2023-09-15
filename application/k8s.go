package application

import (
	"context"
	k8sRepo "kill8s/domain/repository/k8s"
	"kill8s/domain/service"
	"log"
)

var _ k8sRepo.Repository = new(K8sApplication)

type K8sApplication struct {
	logger        *log.Logger
	k8sRepository k8sRepo.Repository
	k8sService    *service.K8sService
}

func NewK8sApplication(l *log.Logger) *K8sApplication {
	ka := &K8sApplication{
		logger:        l,
		k8sRepository: k8sRepo.GetRepository(),
		k8sService:    service.NewK8sService(),
	}
	//ka.logger.SetPrefix("K8sApplication:")

	return ka
}

func (k *K8sApplication) Getter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	res, err := k.k8sRepository.Getter(ctx, namespace, resource, name)
	if err != nil {
		k.logger.Println("Getter err:", err)
	}
	return res, err
}

func (k *K8sApplication) Lister(ctx context.Context, namespace, resource string, opt ...map[string]interface{}) (any, error) {
	res, err := k.k8sRepository.Lister(ctx, namespace, resource)
	if err != nil {
		k.logger.Println("Lister err:", err)
	}
	return res, err
}

func (k *K8sApplication) Deleter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	res, err := k.k8sRepository.Deleter(ctx, namespace, resource, name)
	if err != nil {
		k.logger.Println("Deleter err:", err)
	}
	return res, err
}
