package handlers

import (
	"KubeInsight/internal/server/params"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ks *ServiceHandler) ExportKubeConfig(c *gin.Context) {
	var config params.ClusterConfig
	if err := c.BindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read kubeconfig"})
		return
	}
	if err := ks.kubernetesService.ClusterManagement().ParserKubeConfig([]byte(config.KubeConfig), config.Name, config.ClusterType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kubeconfig parsed successfully"})
}
