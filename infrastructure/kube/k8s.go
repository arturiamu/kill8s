package kube

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var globalK8sClient K8sClient

type K8sClient struct {
	ClientSet     *kubernetes.Clientset
	RESTClient    *rest.RESTClient
	DynamicClient *dynamic.DynamicClient
}

func InitK8s(path *string) {
	if path == nil {
		log.Fatalf("nil path")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *path)
	if err != nil {
		log.Fatalf("Error build k8s config from path: %v", err)
	}
	//dc, err := dynamic.NewForConfig(config)
	//if err != nil {
	//	log.Fatalf("Error init dynamic client from config: %v", err)
	//}
	//globalK8sClient.DynamicClient = dc
	//cs, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	log.Fatalf("Error init client set from config: %v", err)
	//}
	//globalK8sClient.ClientSet = cs

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	rc, err := rest.RESTClientFor(config)
	if err != nil {
		log.Fatalf("Error init rest client from config: %v", err)
	}
	globalK8sClient.RESTClient = rc
}

func GetK8sClient() *K8sClient {
	return &globalK8sClient
}
