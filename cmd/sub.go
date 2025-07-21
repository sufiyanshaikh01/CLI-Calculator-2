package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "sub [num1] [num2]",
	Short: "Subtrack two Number",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		n1, _ := strconv.ParseFloat(args[0], 64)
		n2, _ := strconv.ParseFloat(args[1], 64)
		fmt.Printf("Result:  %.2f\n", n1-n2)
	},
}

func SubCmd() *cobra.Command {
	return subCmd
}
