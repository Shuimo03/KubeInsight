package deployment

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type DeploymentInformer struct {
	Informer informers.SharedInformerFactory
}

func (d *DeploymentInformer) ListResource() (cache.SharedIndexInformer, error) {
	informer := d.Informer.Apps().V1().Deployments().Informer()
	return informer, nil
}
