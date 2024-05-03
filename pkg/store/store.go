package store

import (
	"KubeInsight/pkg/common"
	"KubeInsight/pkg/options"
	"KubeInsight/pkg/store/cache"
	"KubeInsight/pkg/store/mysql"
	"fmt"
	"log"
)

func initRedis() error {
	redisConfig, err := options.LoadRedisConfig("config/store.yaml")
	if err != nil {
		log.Printf("failed to load Redis options: %v", err)
		return err
	}
	dsn := fmt.Sprintf("redis://%s:%s@%s:%d/%d",
		redisConfig.User,
		redisConfig.Password,
		redisConfig.Host,
		redisConfig.Port,
		redisConfig.DB)
	common.Redis, err = cache.NewRedisClient(dsn)
	if err != nil {
		return err
	}
	return nil
}

func initMySQL() error {
	mysqlConfig, err := options.LoadMySQLConfig("config/store.yaml")
	if err != nil {
		log.Printf("failed to load MySQL options: %v", err)
		return err
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
	return nil
}

// TODO 待优化
func InitStoreClient() error {
	if mysqlError := initMySQL(); mysqlError != nil {
		return mysqlError
	}
	if redisError := initRedis(); redisError != nil {
		return redisError
	}
	return nil
}
