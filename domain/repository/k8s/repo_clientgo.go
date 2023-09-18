package k8s

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"kill8s/infrastructure/constant"
	"kill8s/infrastructure/kube"
)

var _ Repository = new(RepoClientGo)

type RepoClientGo struct {
	Client *kube.K8sClient
}

func NewRepoClientGo() RepoClientGo {
	return RepoClientGo{
		Client: kube.GetK8sClient(),
	}
}

func (c RepoClientGo) Creater(ctx context.Context, namespace, resource string, opt ...map[string]interface{}) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (c RepoClientGo) Updater(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (c RepoClientGo) Getter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	var (
		result runtime.Object
		err    error
	)
	if err := constant.Check(resource, constant.ActionGet); err != nil {
		return nil, err
	}
	result = constant.GetRuntimeObj(resource)
	if namespace == "" {
		err = c.Client.RESTClient.Get().Resource(resource).Name(name).Do(ctx).Into(result)
	} else {
		err = c.Client.RESTClient.Get().Namespace(namespace).Resource(resource).Name(name).Do(ctx).Into(result)
	}
	return result, err
}

func (c RepoClientGo) Lister(ctx context.Context, namespace, resource string, opt ...map[string]interface{}) (any, error) {
	var (
		result runtime.Object
		err    error
	)
	if err := constant.Check(resource, constant.ActionGet); err != nil {
		return nil, err
	}
	result = constant.ListRuntimeObj(resource)
	if namespace == "" {
		err = c.Client.RESTClient.Get().Resource(resource).Do(ctx).Into(result)
	} else {
		err = c.Client.RESTClient.Get().Namespace(namespace).Resource(resource).Do(ctx).Into(result)
	}
	err = c.Client.RESTClient.Get().Namespace(namespace).Resource(resource).Do(ctx).Into(result)
	return result, err
}

func (c RepoClientGo) Deleter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	var (
		err error
	)
	if err := constant.Check(resource, constant.ActionGet); err != nil {
		return nil, err
	}
	if namespace == "" {
		err = c.Client.RESTClient.Delete().Resource(resource).Name(name).Do(ctx).Error()
	} else {
		err = c.Client.RESTClient.Delete().Namespace(namespace).Resource(resource).Name(name).Do(ctx).Error()
	}
	return nil, err
}
