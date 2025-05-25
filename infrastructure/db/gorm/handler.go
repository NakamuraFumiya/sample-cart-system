package gorm

import (
	"fmt"
	"sync"

	"github.com/fumiyanakamura/sample-cart-system/infrastructure/config/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	gormHandler *gorm.DB
	once        sync.Once
	initErr     error
)

// スレッドセーフなシングルトンを扱う
func GetInstance() (*gorm.DB, error) {
	once.Do(func() {
		mysqlConfig := viper.LoadConfig().MySQLConfig
		if gormHandler == nil {
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				mysqlConfig.User,
				mysqlConfig.Password,
				mysqlConfig.Host,
				mysqlConfig.Port,
				mysqlConfig.Database,
			)
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				initErr = err
				return
			}
			gormHandler = db
		}
	})

	return gormHandler, initErr
}
