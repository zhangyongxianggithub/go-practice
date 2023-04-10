package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viper.SetConfigName("config") // name of config.toml file (without extension)
	// viper.AddConfigPath("/etc/appname/")  // path to look for the config.toml file in
	// viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")    // optionally look for config.toml in the working directory
	err := viper.ReadInConfig() // Find and read the config.toml file
	if err != nil {             // Handle errors reading the config.toml file
		panic(fmt.Errorf("Fatal error config.toml file: %s \n", err))
	}

}
