package client

import (
	"KubeInsight/model"
	"k8s.io/client-go/kubernetes"
)

func NewKubernetesClient() (kubernetes.Interface, error) {
	kc := model.KubeConfig{}
	k8s, err := kubernetes.NewForConfig(kc.Config)
	if err != nil {
		return nil, err
	}
	return k8s, nil
}
