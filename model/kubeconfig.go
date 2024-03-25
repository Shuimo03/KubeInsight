package model

import "k8s.io/client-go/rest"

type KubeConfig struct {
	Config *rest.Config `json:"config"`
}
