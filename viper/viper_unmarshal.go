package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Format string            `json:"format"`
	ABC    map[string]string `json:"abc"`
}

func main() {
	viper.SetConfigFile("./toml.toml") // 指定配置文件路径
	viper.SetConfigName("config")      // 配置文件名称(无扩展名)
	viper.SetConfigType("toml")        // 如果配置文件的名称中没有扩展名，则需要配置此项
	// viper.AddConfigPath("/Users/zhangyongxiang/Projects/go-practice/viper/") // 查找配置文件所在的路径
	viper.AddConfigPath("./viper") // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()    // 查找并读取配置文件
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			fmt.Println("config.json file not found")
		} else { // 处理读取配置文件的错误
			panic(fmt.Errorf("Fatal error config.json file: %s \n", err))
		}
	}
	var C config

	err = viper.Unmarshal(&C)
	bytes, _ := json.Marshal(C)

	fmt.Println(string(bytes))
}
