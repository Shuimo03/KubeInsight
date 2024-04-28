package informer

import (
	"KubeInsight/pkg/kubernetes/informer/workload/deployment"
	"KubeInsight/pkg/kubernetes/informer/workload/statefulsets"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

/**
node
workload
	Pod
	Statefulset
	Deployment
network
	Service
		lb
		nodePort
		clusterIP
		endpoint
*/

type Interface interface {
	ListResource() (cache.SharedIndexInformer, error)
}

type KubeResource struct {
	Deployment  deployment.DeploymentInformer
	Statefulset statefulsets.StatefulsetInformer
}

func NewKubeResource(factory informers.SharedInformerFactory) *KubeResource {
	return &KubeResource{
		Deployment:  deployment.DeploymentInformer{factory},
		Statefulset: statefulsets.StatefulsetInformer{factory},
	}
}
