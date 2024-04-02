package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewKubernetesClient(kc *rest.Config) (kubernetes.Interface, error) {
	//根据name选择对应kubecConfig
	//判断集群名是否存在,不存在则不需要创建直接退出
	k8s, err := kubernetes.NewForConfig(kc)
	if err != nil {
		return nil, err
	}
	return k8s, nil
}
