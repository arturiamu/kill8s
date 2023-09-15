# kill8s: 轻量的k8s资源管理平台


[English document](https://github.com/arturiamu/kill8s/blob/main/README.md)

# 介绍

kill8s是一个基于gin和client-go的k8s资源管理平台，实现了k8s corev1和k8s appsv1的简单资源的管理。该项目使用领域驱动设计（DDD）架构来组织。

# 使用

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

# 支持的资源类型和操作

| 资源组 | 资源                     | 操作                     |     
|-----|------------------------|------------------------|    
|     | pods                   | delete,get,list        |      
|     | nodes                  | get,list               |     
|     | secrets                | delete,get,list        |   
|     | namespaces             | create,delete,get,list |   
|     | services               | delete,get,list        |   
|     | persistentVolumes      | delete,get,list        |   
|     | replicationcontrollers | delete,get,list        |   
|     | configmaps             | delete,get,list        |   
|     | controllerrevisions    | delete,get,list        |   
|     | daemonsets             | delete,get,list        |   
|     | deployments            | delete,get,list        |   
|     | replicasets            | delete,get,list        |   
|     | statefulsets           | delete,get,list        |

# 未来

- 支持更多的资源和操作
- 基于rbac的权限控制