package k8s

import (
	"context"
)

var globalRepository Repository

type Repository interface {
	ListNamespaces(ctx context.Context) (resp any, err error)
}

func GetRepository() Repository {
	return globalRepository
}

func SetRepository(repo Repository) {
	globalRepository = repo
}
