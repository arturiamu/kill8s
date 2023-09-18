package application

import (
	"context"
	k8sRepo "kill8s/domain/repository/k8s"
	"log"
)

var _ k8sRepo.Repository = new(K8sApplicationMock)

type K8sApplicationMock struct {
	logger *log.Logger
}

func NewK8sApplicationMock(l *log.Logger) *K8sApplicationMock {
	ka := &K8sApplicationMock{
		logger: l,
	}
	return ka
}

type mockResp struct {
	Namespace string `json:"namespace"`
	Resource  string `json:"resource"`
	Name      string `json:"name"`
	Op        string `json:"op"`
}

func (k K8sApplicationMock) Creater(ctx context.Context, namespace, resource string, opt ...map[string]interface{}) (any, error) {
	return mockResp{
		Name:      "",
		Namespace: namespace,
		Resource:  resource,
		Op:        "create",
	}, nil
}

func (k K8sApplicationMock) Deleter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	return mockResp{
		Name:      name,
		Namespace: namespace,
		Resource:  resource,
		Op:        "delete",
	}, nil
}

func (k K8sApplicationMock) Updater(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	return mockResp{
		Name:      name,
		Namespace: namespace,
		Resource:  resource,
		Op:        "update",
	}, nil
}

func (k K8sApplicationMock) Getter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	return mockResp{
		Name:      name,
		Namespace: namespace,
		Resource:  resource,
		Op:        "get",
	}, nil
}

func (k K8sApplicationMock) Lister(ctx context.Context, namespace, resource string, opt ...map[string]interface{}) (any, error) {
	return mockResp{
		Name:      "",
		Namespace: namespace,
		Resource:  resource,
		Op:        "list",
	}, nil
}
