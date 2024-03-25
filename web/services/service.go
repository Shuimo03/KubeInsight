package services

import (
	"KubeInsight/model"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
)

func ParserKubeConfig(kubeConfig []byte) error {
	config, err := clientcmd.RESTConfigFromKubeConfig(kubeConfig)
	if err != nil {
		return fmt.Errorf("Failed to  Parser KubeConfig:%v \n", err)
	}
	_ = model.KubeConfig{ //TODO 持久化存入
		Config: config,
	}
	return nil
}
