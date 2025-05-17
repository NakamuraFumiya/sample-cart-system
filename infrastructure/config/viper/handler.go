package viper

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

func A() {
	viper.SetConfigType("yaml")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	viper.Get("name") // this would be "steve"

	fmt.Println("================")
	fmt.Println(viper.Get("hobbies"))
	fmt.Println("================")

}
