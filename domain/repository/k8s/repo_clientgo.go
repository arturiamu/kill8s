package k8s

import (
	"context"
	"github.com/pkg/errors"
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

func (c RepoClientGo) Getter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	var (
		result runtime.Object
		err    error
	)
	res := constant.NewSupportedResources(resource)
	if !res.ResourceCheck() {
		err = errors.New("unsupported resource type")
		return nil, err
	}
	if !res.ActionCheck(constant.K8sSupportedResourcesActionGet) {
		err = errors.New("unsupported resource action")
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
	res := constant.NewSupportedResources(resource)
	if !res.ResourceCheck() {
		err = errors.New("unsupported resource type")
		return nil, err
	}
	if !res.ActionCheck(constant.K8sSupportedResourcesActionList) {
		err = errors.New("unsupported resource action")
		return nil, err
	}
	result = constant.ListRuntimeObj(resource)
	err = c.Client.RESTClient.Get().Namespace(namespace).Resource(resource).Do(ctx).Into(result)
	return result, err
}

func (c RepoClientGo) Deleter(ctx context.Context, namespace, resource, name string, opt ...map[string]interface{}) (any, error) {
	var (
		err error
	)
	res := constant.NewSupportedResources(resource)
	if !res.ResourceCheck() {
		err = errors.New("unsupported resource type")
		return nil, err
	}
	if !res.ActionCheck(constant.K8sSupportedResourcesActionList) {
		err = errors.New("unsupported resource action")
		return nil, err
	}
	err = c.Client.RESTClient.Get().Namespace(namespace).Resource(resource).Do(ctx).Error()
	return nil, err
}
