package day04

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day04",
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
	matrix := createMatrix(input)
	var count int
	targetWord := "XMAS"

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			for direction := 0; direction < 8; direction++ {
				if isWordValid(matrix, row, col, targetWord, direction) {
					count++
				}
			}
		}
	}
	return count
}

func bonus(input string) int {
	matrix := createMatrix(input)
	var count int
	patterns := []string{"MMSS", "MSSM", "SSMM", "SMMS"}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if string(matrix[row][col]) != "A" {
				continue
			}
			for _, pattern := range patterns {
				if isCrossWordValid(matrix, row, col, pattern) {
					count++
				}
			}
		}
	}
	return count
}

func createMatrix(inputString string) [][]byte {
	lines := strings.Split(inputString, "\r\n")
	matrix := make([][]byte, len(lines))
	for i, line := range lines {
		matrix[i] = []byte(line)
	}
	return matrix
}

func isWordValid(matrix [][]byte, row, col int, targetWord string, direction int) bool {
	dr := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dc := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return false
	}

	for i := 0; i < len(targetWord); i++ {
		newRow := row + i*dr[direction]
		newCol := col + i*dc[direction]

		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) || matrix[newRow][newCol] != targetWord[i] {
			return false
		}
	}
	return true
}

func isCrossWordValid(matrix [][]byte, row, col int, targetWord string) bool {
	dr := []int{-1, -1, 1, 1}
	dc := []int{-1, +1, 1, -1}

	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return false
	}

	for i := 0; i < len(targetWord); i++ {
		newRow := row + dr[i]
		newCol := col + dc[i]

		if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) || matrix[newRow][newCol] != targetWord[i] {
			return false
		}

	}
	return true
}
