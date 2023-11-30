package cmd

import (
	"github.com/spf13/cobra"
)

var sub1Cmd = &cobra.Command{
	Use:   "sub",
	Short: "子命令",
	Long:  `子命令`,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("root persistent pre run, verbose output: ", verbose)
	// },
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	sub1Cmd.Flags().StringVarP(&globalVar, "serverp1", "x", "", "tests")
	rootCmd.AddCommand(sub1Cmd)
}
