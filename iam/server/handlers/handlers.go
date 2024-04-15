package handlers

import (
	"KubeInsight/iam/server/service"
	"KubeInsight/pkg/store/mysql"
)

type IamHandler struct {
	iamService service.IamServiceInterface
}

func NewIamHandler(handler mysql.DB) *IamHandler {
	return &IamHandler{
		service.NewIamService(handler),
	}
}
