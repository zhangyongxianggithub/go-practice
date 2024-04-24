package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("ContentDir", "content")
	fmt.Println(viper.GetString("ContentDir"))

}
