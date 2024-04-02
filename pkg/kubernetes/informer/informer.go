package informer

import (
	"KubeInsight/pkg/kubernetes/informer/workload/deployment"
	"KubeInsight/pkg/kubernetes/informer/workload/statefulsets"
	"k8s.io/client-go/informers"
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
	ListResource() ([]interface{}, error)
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
