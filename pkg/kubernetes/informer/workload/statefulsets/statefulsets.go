package statefulsets

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type StatefulsetInformer struct {
	Informer informers.SharedInformerFactory
}

func (s *StatefulsetInformer) ListResource() (cache.SharedIndexInformer, error) {
	informer := s.Informer.Apps().V1().StatefulSets().Informer()
	return informer, nil
}
