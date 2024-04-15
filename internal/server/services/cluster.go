package services

import (
	"KubeInsight/internal/model"
	"KubeInsight/pkg/store/mysql"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type ClusterManagementInterface interface {
	ParserKubeConfig(kubeConfig []byte, clusterName, clusterType string) error
}

type ClusterManagementService struct {
	dbs mysql.DB
}

func (c *ClusterManagementService) ParserKubeConfig(kubeConfig []byte, clusterName, clusterType string) error {
	config, err := clientcmd.RESTConfigFromKubeConfig(kubeConfig)
	if err != nil {
		return fmt.Errorf("Failed to Parser KubeConfig:%v \n", err.Error())
	}

	cfg := config.String()
	kc := model.KubeConfig{
		Config:    cfg,
		Name:      clusterName,
		Type:      clusterType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	//Panic

	if dbError := c.dbs.GormClient.Create(&kc).Error; dbError != nil {
		return dbError
	}

	return nil
}

func newClusterManagementService(srv *Services) *ClusterManagementService {
	return &ClusterManagementService{dbs: srv.dbHandler}
}
