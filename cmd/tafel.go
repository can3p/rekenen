/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

func TafelCommand() *cobra.Command {
	var valsStr string
	var number int64

	var tafelCmd = &cobra.Command{
		Use:   "tafel",
		Short: "Generated multiplication tables",
		RunE: func(cmd *cobra.Command, args []string) error {
			vals := strings.Split(valsStr, ",")

			for number > 0 {
				v1 := vals[rand.Intn(len(vals))]
				v2 := 1 + rand.Int63n(9)

				fmt.Printf("%-2s x %-2d = ___\n", v1, v2)
				number--
			}

			return nil
		},
	}

	tafelCmd.Flags().StringVarP(&valsStr, "vals", "v", "1,2,3,4,5,6,7,8,9", "vals")
	tafelCmd.Flags().Int64VarP(&number, "number", "n", 10, "number value")

	return tafelCmd
} // tafelCmd represents the tafel command

func init() {
	rootCmd.AddCommand(TafelCommand())
}
