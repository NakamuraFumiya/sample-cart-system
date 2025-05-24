package gorm

import (
	"fmt"

	"github.com/fumiyanakamura/sample-cart-system/infrastructure/config/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormHandler *gorm.DB

// FIXME: スレッドセーフになっていないので修正する
func GetInstance() (*gorm.DB, error) {
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
			return nil, err
		}
		gormHandler = db
	}
	return gormHandler, nil
}
