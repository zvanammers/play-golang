package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wc-tool",
	Long:  `All software has versions. This is wc-tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wc-tool v0.1")
	},
}
