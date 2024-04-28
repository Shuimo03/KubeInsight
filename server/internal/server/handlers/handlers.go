package handlers

import (
	"KubeInsight/pkg/store/mysql"
	"KubeInsight/server/internal/server/services"
)

type ServiceHandler struct {
	kubernetesService services.ServiceInterface
}

func NewServiceHandlerHandler(handler mysql.DB) *ServiceHandler {
	return &ServiceHandler{
		kubernetesService: services.NewServices(handler),
	}
}
