package handlers

import (
	"KubeInsight/internal/server/services"
	"KubeInsight/pkg/store/mysql"
)

type ServiceHandler struct {
	kubernetesService services.ServiceInterface
}

func NewServiceHandlerHandler(handler mysql.DB) *ServiceHandler {
	return &ServiceHandler{
		kubernetesService: services.NewServices(handler),
	}
}
