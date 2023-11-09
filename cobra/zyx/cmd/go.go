package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "go 子命令",
	Long:  `go 子命令`,
	Run: func(cmd *cobra.Command, args []string) {
		// 添加这个，不添加任何参数，会执行这个函数，如果不允许执行go，去掉这个函数
	},
}

func createGoCommand() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "create",
		Short: "create go",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("command: ", cmd.Name(), cmd.Short, cmd.Long)
			fmt.Println("verbose value: ", verbose)
			fmt.Println("args: ", args, name, globalVar)
		},
	}
	command.PersistentFlags().StringVarP(&name, "name", "n", "", "name名字")
	_ = viper.BindPFlag("gbi.go.name", command.PersistentFlags().Lookup("name"))
	return command
}
func getGoCommand() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "get",
		Short: "get go",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("command: ", cmd.Name(), cmd.Short, cmd.Long)
			fmt.Println("args: ", args)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name名字")
	return command
}
func init() {
	goCmd.AddCommand(createGoCommand())
	goCmd.AddCommand(getGoCommand())
	rootCmd.AddCommand(goCmd)
}
