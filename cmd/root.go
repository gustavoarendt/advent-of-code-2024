package cmd

import (
	"fmt"
	"os"

	"github.com/gustavoarendt/advent-of-code-2024/cmd/2024/day01"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2024",
	Short: "Advent of Code",
	Long:  "Advent of Code 2024 - Golang version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	Cmd.AddCommand(day01.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Help() {
	print("Uses go run main.go 'year' 'day' to run the program\n")
}
