package day01

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day01",
	Short: "Run first day puzzle",
	Long:  "Run first day puzzle",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	content, err := os.ReadFile(fmt.Sprintf(`cmd/%s/%s/input.txt`, parent, command))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 01 - Puzzle: %d\n", puzzle(string(content)))
	fmt.Printf("Day 01 - Bonus: %d\n", bonus(string(content)))
}

func puzzle(input string) int {
	if input == "" {
		return 0
	}
	return 1
}

func bonus(input string) int {
	if input == "" {
		return 0
	}
	return 1
}
