package service

import (
	"KubeInsight/pkg/common"
	"KubeInsight/pkg/store/mysql"
	"github.com/redis/go-redis/v9"
)

type IamServiceInterface interface {
	Login(username, password string) error
	Auth(userName string) error
}

type IamService struct {
	dbHandler    *mysql.DB
	cacheHandler *redis.Client
}

func NewIamService() *IamService {
	service := &IamService{
		common.DB,
		common.Redis,
	}
	return service
}
