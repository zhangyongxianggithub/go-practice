package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.RegisterAlias("loud", "Verbose") // 注册别名（此处loud和Verbose建立了别名）

	viper.Set("verbose", true) // 结果与下一行相同
	viper.Set("loud", false)   // 结果与前一行相同

	fmt.Println(viper.GetBool("loud"))    // true
	fmt.Println(viper.GetBool("verbose")) // true

}
