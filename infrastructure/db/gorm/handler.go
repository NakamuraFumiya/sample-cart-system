package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormHandler *gorm.DB

// FIXME: スレッドセーフになっていないので修正する
// FIXME: 環境変数から読めるようにする
func GetInstance() (*gorm.DB, error) {
	if gormHandler == nil {
		dsn := "development_user:development_password@tcp(127.0.0.1:3306)/development_database?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		gormHandler = db
	}
	return gormHandler, nil
}
