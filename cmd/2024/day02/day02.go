package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day02",
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
	var safeCount int
	for _, line := range strings.Split(input, "\r\n") {
		numbers := convertStringIntoNumbers(strings.Fields(line))
		if isSafe(numbers) {
			safeCount++
		}
	}
	return safeCount
}

func bonus(input string) int {
	var safeCount int
	for _, line := range strings.Split(input, "\r\n") {
		numbers := convertStringIntoNumbers(strings.Fields(line))
		if isSafe(numbers) {
			safeCount++
			continue
		}

		for idx := range numbers {
			reCheck := append(append([]int{}, numbers[:idx]...), numbers[idx+1:]...)
			if isSafe(reCheck) {
				safeCount++
				break
			}
		}
	}
	return safeCount
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func convertStringIntoNumbers(array []string) []int {
	var intArray = []int{}
	for _, val := range array {
		number, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		intArray = append(intArray, number)
	}
	return intArray
}

func isSafe(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	order := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if order*diff <= 0 || abs(diff) > 3 {
			return false
		}
	}
	return true
}
