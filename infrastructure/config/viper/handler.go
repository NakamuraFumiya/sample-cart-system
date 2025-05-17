package viper

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println("================")
	fmt.Println(viper.Get("rdb"))
	fmt.Println("================")
}
