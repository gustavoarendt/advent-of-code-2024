package day05

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day05",
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
	var total int
	rules := parseRules(input)
	updates := parseUpdates(input)

	for _, update := range updates {
		if isCorrectOrder(update, rules) {
			total += findInnerPage(update)
		}
	}
	return total
}

func bonus(input string) int {
	var total int
	rules := parseRules(input)
	updates := parseUpdates(input)

	for _, update := range updates {
		if !isCorrectOrder(update, rules) {
			update = reArrange(update, rules)
			total += findInnerPage(update)
		}
	}
	return total
}

func reArrange(update []int, rules map[int][]int) []int {
	sortedUpdate := make([]int, len(update))
	copy(sortedUpdate, update)

	for i := 0; i < len(sortedUpdate)-1; i++ {
		for j := i + 1; j < len(sortedUpdate); j++ {
			if _, ok := rules[sortedUpdate[i]]; ok {
				for _, rulePage := range rules[sortedUpdate[i]] {
					if rulePage == sortedUpdate[j] {
						sortedUpdate[i], sortedUpdate[j] = sortedUpdate[j], sortedUpdate[i]
						break
					}
				}
			}
		}
	}
	return sortedUpdate
}

func findInnerPage(update []int) int {
	return update[len(update)/2]
}

func isCorrectOrder(update []int, rules map[int][]int) bool {
	for i := len(update) - 1; i >= 0; i-- {
		if _, ok := rules[update[i]]; ok {
			for _, rulePage := range rules[update[i]] {
				for _, remaining := range update[:i] {
					if rulePage == remaining {
						return false
					}
				}
			}
		}
	}
	return true
}

func parseRules(input string) map[int][]int {
	rules := make(map[int][]int)
	for _, line := range strings.Split(input, "\r\n") {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			rules[a] = append(rules[a], b)
		}
	}
	return rules
}

func parseUpdates(input string) [][]int {
	var updates [][]int
	for _, line := range strings.Split(input, "\r\n") {
		if strings.Contains(line, ",") {
			update := []int{}
			for _, page := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(page)
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}
	return updates
}
