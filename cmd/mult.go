/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

func MultCommand() *cobra.Command {
	var minVal int64
	var maxVal int64
	var number int64

	var multCmd = &cobra.Command{
		Use:   "mult",
		Short: "Generated addition examples",
		RunE: func(cmd *cobra.Command, args []string) error {
			for number > 0 {
				v1 := minVal + rand.Int63n(maxVal-minVal+1)
				v2 := minVal + rand.Int63n(maxVal-minVal+1)

				fmt.Printf("%-2d x %-2d = ___\n", v1, v2)
				number--
			}

			return nil
		},
	}

	multCmd.Flags().Int64VarP(&minVal, "min", "f", 1, "min value")
	multCmd.Flags().Int64VarP(&maxVal, "max", "t", 10, "max value")
	multCmd.Flags().Int64VarP(&number, "number", "n", 10, "number value")

	return multCmd
} // multCmd represents the mult command

func init() {
	rootCmd.AddCommand(MultCommand())
}
