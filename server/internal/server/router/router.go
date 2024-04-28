package router

import (
	"KubeInsight/pkg/common"
	"KubeInsight/server/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	kubernetesGroup := r.Group("/v1/kubernetes")
	kc := handlers.NewServiceHandlerHandler(*common.DB)
	kubernetesGroup.POST("/cluster/config", kc.ExportKubeConfig)
	return r
}
