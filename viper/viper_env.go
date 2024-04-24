package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	fmt.Println(os.Environ())
	viper.AutomaticEnv()
	viper.SetEnvPrefix("HOMEBREW") // 将自动转为大写
	err := viper.BindEnv("PREFIX")
	if err != nil {
		return
	}
	id := viper.Get("PREFIX") // 13
	fmt.Println(id)

}
