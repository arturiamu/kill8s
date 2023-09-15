# kill8s: Lightweight k8s resource management platform

[中文文档](https://github.com/arturiamu/kill8s/blob/main/README.ZH.md)

# Introduction

kill8s is a k8s resource management platform based on gin and client-go.
It realizes the management of simple resources of k8s corev1 and k8s appsv1. 
The project is organized using Domain-driven design (DDD) architecture.

# Usage

```bash
git clone https://github.com/arturiamu/kill8s.git
cd kill8s
go run main.go -kubeconfig=kubeconfigpath
```

```bash
curl http://localhost:8888/api/v1/k8s/api/v1/nodes/minikube
curl http://localhost:8888/api/v1/k8s/api/v1/namespaces/kube-system/pods
curl http://localhost:8888/api/v1/k8s/api/v1/namespaces/kube-system/pods/etcd-minikube
```

# Supported resources and operations

| group | resource               | operations             |     
|-------|------------------------|------------------------|    
|       | pods                   | delete,get,list        |      
|       | nodes                  | get,list               |     
|       | secrets                | delete,get,list        |   
|       | namespaces             | create,delete,get,list |   
|       | services               | delete,get,list        |   
|       | persistentVolumes      | delete,get,list        |   
|       | replicationcontrollers | delete,get,list        |   
|       | configmaps             | delete,get,list        |   
|       | controllerrevisions    | delete,get,list        |   
|       | daemonsets             | delete,get,list        |   
|       | deployments            | delete,get,list        |   
|       | replicasets            | delete,get,list        |   
|       | statefulsets           | delete,get,list        |
       
# Future

- Support more resource types and operations
- Permission control based on rbac