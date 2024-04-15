package services

import (
	"KubeInsight/pkg/store/mysql"
)

type ServiceInterface interface {
	ClusterManagement() ClusterManagementInterface
}

type Services struct {
	dbHandler mysql.DB
}

func (s *Services) ClusterManagement() ClusterManagementInterface {
	return newClusterManagementService(s)
}

func NewServices(handler mysql.DB) *Services {
	return &Services{
		dbHandler: handler,
	}
}
