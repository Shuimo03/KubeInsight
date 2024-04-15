package service

import "KubeInsight/pkg/store/mysql"

type IamServiceInterface interface {
	Login(username, password string) error
}

type IamService struct {
	dbHandler mysql.DB
}

func NewIamService(dbHandler mysql.DB) *IamService {
	service := &IamService{
		dbHandler,
	}
	return service
}
