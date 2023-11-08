package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/gookit/goutil/fsutil"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CommandName() string {
	commandPath, err := os.Executable()
	if err != nil {
		commandPath = "unknown"
	}
	command := filepath.Base(commandPath)
	return command
}

// rootCmd 代表没有调用子命令时的基础命令
var rootCmd = &cobra.Command{
	Use:   CommandName(),
	Short: "一个简单的测试命令",
	Long:  `一个简单的测试命令`,
}

// Execute 将所有子命令添加到root命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 调用一次。
func Execute() {
	initConfig()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	configName := CommandName()
	configPath := os.Getenv("HOME") + "/." + strings.ToLower(configName) + "/config.yaml"
	configEnvPrefix := strings.ToUpper(configName)

	if fsutil.FileExists(configPath) {
		viper.SetConfigFile(configPath)
	}
	viper.SetEnvPrefix(configEnvPrefix)

	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("CommandConfigInitError read in: %v", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("CommandConfigChanged: %s", e.Name)
	})
}
