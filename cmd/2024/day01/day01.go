package day01

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

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
	var score int
	leftList := []int{}
	rightList := []int{}
	for _, line := range strings.Split(input, "\r\n") {
		numbers := strings.Fields(line)
		populateSlice(numbers, &leftList, &rightList)
	}
	sort.Ints(leftList)
	sort.Ints(rightList)
	for i := 0; i < len(leftList); i++ {
		score += abs(leftList[i] - rightList[i])
	}
	return score
}

func bonus(input string) int {
	var similarity int
	leftList := []int{}
	rightList := []int{}
	for _, line := range strings.Split(input, "\r\n") {
		numbers := strings.Fields(line)
		populateSlice(numbers, &leftList, &rightList)
	}
	for _, lVal := range leftList {
		acc := 0
		for _, rVal := range rightList {
			if rVal == lVal {
				acc++
			}
		}
		similarity += acc * lVal
	}
	return similarity
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func insert(slice *[]int, value string) []int {
	convertedValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Errorf("error converting %s to integer: %w", value, err))
	}
	return append(*slice, convertedValue)
}

func populateSlice(numbers []string, sliceEven, sliceOdd *[]int) {
	count := 0
	for _, val := range numbers {
		if count%2 == 0 {
			*sliceEven = insert(sliceEven, val)
		} else {
			*sliceOdd = insert(sliceOdd, val)
		}
		count++
	}
}
