package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml") // 或者 viper.SetConfigType("YAML")

	// 任何需要将此配置添加到程序中的方法。
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

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		return
	}
	viper.Set("name", "setting it")

	fmt.Println(viper.Get("name")) // 这里会得到 "steve"

}
