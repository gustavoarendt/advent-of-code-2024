package cmd

import (
	"fmt"
	"os"

	"github.com/gustavoarendt/advent-of-code-2024/cmd/2024/day01"
	"github.com/gustavoarendt/advent-of-code-2024/cmd/2024/day02"
	"github.com/gustavoarendt/advent-of-code-2024/cmd/2024/day03"
	"github.com/gustavoarendt/advent-of-code-2024/cmd/2024/day04"
	"github.com/gustavoarendt/advent-of-code-2024/cmd/2024/day05"
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
	Cmd.AddCommand(day02.Cmd)
	Cmd.AddCommand(day03.Cmd)
	Cmd.AddCommand(day04.Cmd)
	Cmd.AddCommand(day05.Cmd)
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
