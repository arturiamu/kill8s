package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var clientSet *kubernetes.Clientset

func InitK8s(path *string) {
	if path == nil {
		log.Fatalf("nil path")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *path)
	if err != nil {
		log.Fatalf("Error build k8s config from path: %v", err)
	}
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error init clientSet from config: %v", err)
	}
	clientSet = cs
}

func GetK8sClientSet() *kubernetes.Clientset {
	return clientSet
}
