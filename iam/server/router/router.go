package router

import (
	"KubeInsight/iam/server/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	iam := handlers.NewIamHandler()

	kubernetesGroup := r.Group("/v1/iam")
	kubernetesGroup.POST("/login", iam.Login)
	kubernetesGroup.POST("/auth", iam.Auth)
	return r
}
