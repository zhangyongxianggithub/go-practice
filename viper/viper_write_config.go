package main

import "github.com/spf13/viper"

func main() { // 指定配置文件路径
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("json")
	viper.AddConfigPath("./viper") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")
	viper.Set("format", "json-rewrite")
	viper.Set("format1", "json-rewrite")
	// viper.WriteConfig() // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
	viper.SafeWriteConfig()
	// viper.WriteConfigAs("./config-as")
	//viper.SafeWriteConfigAs("./config-as") // 因为该配置文件写入过，所以会报错
	//viper.SafeWriteConfigAs("./config-as-other")
}
