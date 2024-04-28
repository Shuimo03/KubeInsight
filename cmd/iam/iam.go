package main

import (
	"KubeInsight/iam/server/router"
	"KubeInsight/pkg/common"
	"KubeInsight/pkg/options"
	"KubeInsight/pkg/store/mysql"
	"KubeInsight/pkg/store/redis"
	"context"
	"fmt"
	"log"
)

// 抽象成公共函数
func initMySQLClient() {
	mysqlConfig, err := options.LoadMySQLConfig("config/store.yaml")

	if err != nil {
		log.Fatalf("failed to load MySQL options: %v", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConfig.DataBase.Username,
		mysqlConfig.DataBase.Password,
		mysqlConfig.DataBase.Host,
		mysqlConfig.DataBase.Port,
		mysqlConfig.DataBase.Database,
		mysqlConfig.DataBase.Charset,
		mysqlConfig.DataBase.ParseTime,
		mysqlConfig.DataBase.Loc)
	common.DB, err = mysql.NewMySQLClient(dsn)

	//redis://user:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2

}

func initRedis() {
	redisConfig, err := options.LoadRedisConfig("config/store.yaml")
	redisURL := fmt.Sprintf("redis://user:password@%s:%s", redisConfig.Host, redisConfig.Port)
	common.Redis, err = redis.NewRedisClient(redisURL)
	common.Redis.Set(context.TODO(), "Test", "T", 0)
	if err != nil {
		log.Fatalf("mysql init failed: %v", err)
	}
}

func init() {
	initMySQLClient()
	//if err := model.InitModel(); err != nil {
	//	panic(err)
	//}
	//config.InitSySAdmin()
	initRedis()
}

func main() {
	r := router.Router()
	r.Run()
}
