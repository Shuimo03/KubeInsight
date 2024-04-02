package statefulsets

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type StatefulsetInformer struct {
	Informer informers.SharedInformerFactory
}

func (s *StatefulsetInformer) ListResource() ([]interface{}, error) {
	informer := s.Informer.Apps().V1().StatefulSets().Informer()
	if !cache.WaitForCacheSync(nil, informer.HasSynced) {
		return nil, fmt.Errorf("timed out waiting for caches to sync")
	}
	deployments := informer.GetStore().List()
	return deployments, nil
}
