package k8s

import "context"

var globalRepository Repository

type Repository interface {
	Creater(ctx context.Context, namespace, resource string, opt ...map[string]interface{}) (any, error)
	Deleter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error)
	Updater(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error)
	Getter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error)
	Lister(ctx context.Context, namespace, resource string, opt ...map[string]interface{}) (any, error)
}

func GetRepository() Repository {
	return globalRepository
}

func SetRepository(repo Repository) {
	globalRepository = repo
}
