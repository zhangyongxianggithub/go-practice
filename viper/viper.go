package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viper.SetConfigName("config.json") // name of toml.toml file (without extension)
	// viper.AddConfigPath("/etc/appname/")  // path to look for the toml.toml file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for toml.toml in the working directory
	err := viper.ReadInConfig() // Find and read the toml.toml file
	if err != nil {             // Handle errors reading the toml.toml file
		panic(fmt.Errorf("Fatal error toml.toml file: %s \n", err))
	}
	fmt.Println(viper.GetString("ContentDir"))

}
