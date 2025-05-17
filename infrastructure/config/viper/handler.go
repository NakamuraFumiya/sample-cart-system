package viper

import (
	"encoding/json"
	"fmt"

	"github.com/fumiyanakamura/sample-cart-system/infrastructure/config"
	"github.com/spf13/viper"
)

func Load() config.Config {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config.Config{
		MySQLConfig: BindMySQLConfig(),
	}
}

func BindMySQLConfig() config.MySQLConfig {
	// map → JSON
	// JSONにエンコードする
	jsonData, err := json.Marshal(viper.Get("rdb.mysql"))
	if err != nil {
		panic(fmt.Errorf("json marshal error: %w", err))
	}

	// JSON → struct
	// 構造体にデコードする
	var mysqlConfig config.MySQLConfig
	if err := json.Unmarshal(jsonData, &mysqlConfig); err != nil {
		panic(fmt.Errorf("json unmarshal error: %w", err))
	}

	return mysqlConfig
}
