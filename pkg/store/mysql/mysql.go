package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

//TODO 需要初始化所有model到库中

var (
	db      *gorm.DB
	dberror error
	once    sync.Once
)

type DB struct {
	GormClient *gorm.DB
}

func NewMySQLClient(dsn string) (*DB, error) {
	once.Do(func() {
		db, dberror = gorm.Open(mysql.New(mysql.Config{
			DSN: dsn, // "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local" data source name, 详情参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
		}), &gorm.Config{})

		if dberror != nil {
			log.Println("Failed init MySQL:", dberror)
		}
	})

	d := &DB{
		db,
	}

	return d, nil
}
