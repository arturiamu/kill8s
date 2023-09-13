package application

import (
	k8sRepo "kill8s/domain/repository/k8s"
	"kill8s/domain/service"
	"log"
)

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
	ka.logger.SetPrefix("K8sApplication:")

	return ka
}
