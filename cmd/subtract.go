package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var subtractCMD = &cobra.Command{
	Use:   "subtract",
	Short: "Subtract one integer from another",
	Long:  `Subtract one integer a from another integer b: result = a - b`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		var err error
		a, err = strconv.Atoi(args[0])
		if err != nil {
			panic("Arguments to `subtract` must be integers")
		}
		b, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Arguments to `subtract` must be integers")
		}

		result := a - b
		fmt.Println(result)
	},
}

func init() {
	rootCMD.AddCommand(subtractCMD)
}
