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
	Run:   func(cmd *cobra.Command, args []string) {},
}

func createGoCommand() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "create",
		Short: "create go",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("command: ", cmd.Name(), cmd.Short, cmd.Long)
			fmt.Println("args: ", args, name)
			fmt.Println("verbose value: ", verbose)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name名字")
	_ = viper.BindPFlag("gbi.go.name", command.Flags().Lookup("name"))
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
