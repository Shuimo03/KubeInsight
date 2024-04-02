package deployment

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type DeploymentInformer struct {
	Informer informers.SharedInformerFactory
}

func (d *DeploymentInformer) ListResource(stopCh chan struct{}) ([]interface{}, error) {
	informer := d.Informer.Apps().V1().Deployments().Informer()
	if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
		return nil, fmt.Errorf("timed out waiting for caches to sync")
	}
	deployments := informer.GetStore().List()
	return deployments, nil
}
