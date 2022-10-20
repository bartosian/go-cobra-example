package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var addCMD = &cobra.Command{
	Use:   "add",
	Short: "Add two integers",
	Long:  `Add two integers a and b; result = a + b`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		var err error

		a, err = strconv.Atoi(args[0])
		if err != nil {
			panic("Arguments to `add` must be integers")
		}

		b, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Arguments to `add` must be integers")
		}

		result := a + b

		fmt.Println(result)
	},
}

func init() {
	rootCMD.AddCommand(addCMD)
}
