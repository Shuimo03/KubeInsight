package main

import (
	"KubeInsight/pkg/store/mysql"
	"KubeInsight/server/internal/common"
	"KubeInsight/server/internal/model"
	"KubeInsight/server/internal/options"
	"KubeInsight/server/internal/server/router"
	"fmt"
	"log"
)

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
	if err != nil {
		log.Fatalf("mysql init failed: %v", err)
	}
	kubeconfig := model.KubeConfig{}
	if autoMigrate := common.DB.GormClient.AutoMigrate(&kubeconfig); autoMigrate != nil {
		log.Fatalf("Table init: %v", autoMigrate)
	}

}

func init() {
	initMySQLClient()
}

func main() {
	r := router.Router()
	r.Run()
}
