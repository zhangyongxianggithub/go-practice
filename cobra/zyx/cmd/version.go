package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "go",
	Short: "go 子命令",
	Long:  `go 子命令`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version")
		fmt.Println("command: ", cmd.Name(), cmd.Short, cmd.Long)
		fmt.Println("args: ", args)
	},
}

func createCommand() *cobra.Command {
	var arg string
	command := &cobra.Command{
		Use:   "create",
		Short: "create go",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version")
			fmt.Println("command: ", cmd.Name(), cmd.Short, cmd.Long)
			fmt.Println("args: ", args)
		},
	}
	command.Flags().StringVar(&arg, "args", "", "Workspace id")
	return command
}
func init() {
	versionCmd.AddCommand(createCommand())
	rootCmd.AddCommand(versionCmd)
}
