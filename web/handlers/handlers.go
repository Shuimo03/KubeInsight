package handlers

import (
	"KubeInsight/web/params"
	"KubeInsight/web/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExportKubeConfig(c *gin.Context) {
	var config params.KubeConfig
	if err := c.BindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON"})
		return
	}
	err := services.ParserKubeConfig([]byte(""))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User received successfully"})
}
