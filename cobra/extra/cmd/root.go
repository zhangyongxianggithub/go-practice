package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CommandName() string {
	commandPath, _ := os.Executable()
	return filepath.Base(commandPath)
}

// rootCmd 代表没有调用子命令时的基础命令
var rootCmd = &cobra.Command{
	Use:                        CommandName(),
	Short:                      "一个简单的测试命令",
	Long:                       `一个简单的测试命令`,
	Version:                    "1.0",
	SuggestionsMinimumDistance: 3,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("root persistent pre run, verbose output: ", verbose)
	// },
}

var verbose bool

// Execute 将所有子命令添加到root命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 调用一次。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var globalVar string
var buildDate = time.Now().Format("20060102150405")

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate(CommandName() + " version " + rootCmd.Version + " build date " + buildDate + "\n")
	rootCmd.SetUsageFunc(func(command *cobra.Command) error {

	})
	rootCmd.PersistentFlags().StringVarP(&globalVar, "global", "g", "", "tests")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
}

func initConfig() {
	configName := CommandName()
	configEnvPrefix := strings.ToUpper(configName)
	home, err := homedir.Dir()
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err == nil {
		viper.AddConfigPath(home)
		viper.SetConfigName(".zyx_config")
	}
	if env, ok := os.LookupEnv(configEnvPrefix); ok {
		viper.SetEnvPrefix(env)
	}
	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		logger.Errorf("read config failed: %v", err)
	}
	fmt.Println("init config run end")
	fmt.Println("key::::", viper.GetString("gbi.go.name"))
}
