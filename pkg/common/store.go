package common

import (
	"KubeInsight/pkg/store/mysql"
	"github.com/redis/go-redis/v9"
)

var DB *mysql.DB
var Redis *redis.Client
