package router

import (
	"KubeInsight/web/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	kubernetesGroup := r.Group("/v1/kubernetes")
	kubernetesGroup.POST("/config", handlers.ExportKubeConfig)
	return r
}
