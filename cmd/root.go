package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCMD = &cobra.Command{
	Use:   "calc",
	Short: "Calculate arithmetic expressions",
	Long: `Calculate arithmetic expressions.

           It can add integers and subtract integers.`,
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}