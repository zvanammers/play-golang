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
	Short: "Print the version number of basic-cli",
	Long:  `All software has versions. This is basic-clis`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("basic-cli v0.1 -- HEAD")
	},
}
