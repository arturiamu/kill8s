package kube

import (
	"context"
	"fmt"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	var kubeconfig = "/Users/mulinbiao/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error build k8s config from path: %v", err)
	}
	initApi(config)
}

func initApi(config *rest.Config) {
	//config.APIPath = "api"
	//config.GroupVersion = &corev1.SchemeGroupVersion
	config.APIPath = "apis"
	config.GroupVersion = &batchv1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	RESTClient, err := rest.RESTClientFor(config)
	if err != nil {
		log.Fatalf("Error init rest client from config: %v", err)
	}

	err = RESTClient.Get().Resource("replicasets").Do(context.Background()).Error()
	fmt.Printf("get err:%v\n", err)
}
