package repository

import (
	k8sRepo "kill8s/domain/repository/k8s"
	userRepo "kill8s/domain/repository/user"
)

func Init() {
	userRepo.SetRepository(userRepo.NewRepoMysql())
	k8sRepo.SetRepository(k8sRepo.NewRepoClientGo())
}
