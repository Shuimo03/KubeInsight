package router

import (
	"KubeInsight/iam/server/handlers"
	"KubeInsight/pkg/common"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	iam := handlers.NewIamHandler(*common.DB)

	kubernetesGroup := r.Group("/v1/iam")
	kubernetesGroup.POST("/login", iam.Login)
	return r
}
