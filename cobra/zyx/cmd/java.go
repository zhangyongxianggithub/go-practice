package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "java 子命令",
	Long:  `java 子命令`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func createJavaCommand() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "create",
		Short: "create java",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("command: ", cmd.Name(), cmd.Short, cmd.Long)
			fmt.Println("args: ", args)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name名字")
	return command
}
func getJavaCommand() *cobra.Command {
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
	javaCmd.AddCommand(createJavaCommand())
	javaCmd.AddCommand(getJavaCommand())
	rootCmd.AddCommand(javaCmd)
}
