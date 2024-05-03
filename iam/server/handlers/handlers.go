package handlers

import (
	"KubeInsight/iam/server/service"
	"github.com/gin-gonic/gin"
)

type IamHandler struct {
	iamService service.IamServiceInterface
}

type IamHandlerInterface interface {
	Login(c *gin.Context)
	Auth(c *gin.Context)
}

func NewIamHandler() IamHandlerInterface {
	return &IamHandler{
		service.NewIamService(),
	}
}
