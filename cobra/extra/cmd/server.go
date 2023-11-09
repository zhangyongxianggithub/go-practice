package cmd

import (
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "server",
	Short: "子命令",
	Long:  `子命令`,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("root persistent pre run, verbose output: ", verbose)
	// },
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	subCmd.Flags().StringVarP(&globalVar, "serverp", "g", "", "tests")
	rootCmd.AddCommand(subCmd)
}
