package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(additionCmd)
}

var additionCmd = &cobra.Command{
	Use:   "addition",
	Short: "Add two numbers",
	Long:  `Add two positive natural numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Wrong number of provided args. Number of args must be 2.")
			return
		}
		num1, num1Err := strconv.Atoi(args[0])
		num2, num2Err := strconv.Atoi(args[1])
		if num1Err != nil || num2Err != nil {
			if num1Err != nil {
				fmt.Printf("Error converting the first argument -> %s\n", num1Err)
			}
			if num2Err != nil {
				fmt.Printf("Error converting the second argument -> %s\n", num2Err)
			}
			return
		}
		fmt.Printf("The sum of %d and %d is %d.\n", num1, num2, num1+num2)
	},
}
