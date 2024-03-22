/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

func PlusCommand() *cobra.Command {
	var minVal int64
	var maxVal int64
	var number int64

	var plusCmd = &cobra.Command{
		Use:   "plus",
		Short: "Generated addition examples",
		RunE: func(cmd *cobra.Command, args []string) error {
			for number > 0 {
				v1 := minVal + rand.Int63n(maxVal-minVal+1)
				v2 := minVal + rand.Int63n(maxVal-minVal+1)

				fmt.Printf("%-2d + %-2d = ___\n", v1, v2)
				number--
			}

			return nil
		},
	}

	plusCmd.Flags().Int64VarP(&minVal, "min", "f", 1, "min value")
	plusCmd.Flags().Int64VarP(&maxVal, "max", "t", 10, "max value")
	plusCmd.Flags().Int64VarP(&number, "number", "n", 10, "number value")

	return plusCmd
} // plusCmd represents the plus command

func init() {
	rootCmd.AddCommand(PlusCommand())
}
