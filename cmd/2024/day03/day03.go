package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day03",
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
	return filterAndSumMatches(input)
}

func bonus(input string) int {
	var totalWithConditions int
	splittedByDont := strings.Split(input, "don't()")
	totalWithConditions += filterAndSumMatches(splittedByDont[0])
	for _, val := range splittedByDont {
		index := strings.Index(val, "do()")
		if index != -1 {
			result := val[index:]
			totalWithConditions += filterAndSumMatches(result)
		}
	}
	return totalWithConditions
}

func multiplyMul(str string) int {
	substring := strings.Split(strings.Replace(strings.Replace(str, ")", "", 1), "mul(", "", 1), ",")
	params := convertStringIntoNumbers(substring)
	return params[0] * params[1]
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

func filterAndSumMatches(input string) int {
	var total int
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllString(input, -1)
	for _, match := range matches {
		total += multiplyMul(match)
	}
	return total
}
