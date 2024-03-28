package main

import (
	"KubeInsight/internal/common"
	"KubeInsight/internal/options"
	"KubeInsight/internal/web/router"
	"KubeInsight/pkg/store/mysql"
	"fmt"
	"log"
)

func initMySQLClient() {
	mysqlConfig, err := options.LoadMySQLConfig("config/mysql.yaml")
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

}

func init() {
	initMySQLClient()
}

func main() {
	r := router.Router()
	r.Run()
}
