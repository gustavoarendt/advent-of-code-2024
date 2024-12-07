package day06

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day06",
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
	var count int
	matrix := createMatrix(input)
	sRow, sCol := findInitialPoint(matrix)
	plan := trackGuard(matrix, sRow, sCol)
	for i := range plan {
		for j := range plan[0] {
			// print(string(plan[i][j]))
			if plan[i][j] == 'X' {
				count++
			}
		}
		// print("\n")
	}
	return count
}

func bonus(input string) int {
	matrix := createMatrix(input)
	sRow, sCol := findInitialPoint(matrix)
	count := trackGuardAndTrap(matrix, sRow, sCol)
	return count
}

func findInitialPoint(matrix [][]byte) (row, col int) {
	var sRow int
	var sCol int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == '^' {
				sRow = row
				sCol = col
				matrix[sRow][sCol] = 'X'
			}
		}
	}
	return sRow, sCol
}

func createMatrix(inputString string) [][]byte {
	lines := strings.Split(inputString, "\r\n")
	matrix := make([][]byte, len(lines))
	for i, line := range lines {
		matrix[i] = []byte(line)
	}
	return matrix
}

func trackGuard(matrix [][]byte, sRow, sCol int) [][]byte {
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}
	direction := 0
	matrix = move(matrix, sRow, sCol, dr, dc, direction)
	return matrix
}

func move(matrix [][]byte, sRow, sCol int, dr, dc []int, direction int) [][]byte {
	row := sRow + dr[direction]
	col := sCol + dc[direction]
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return matrix
	}
	if matrix[row][col] == '.' || matrix[row][col] == 'X' {
		matrix[row][col] = 'X'
		return move(matrix, row, col, dr, dc, direction)
	}
	if matrix[row][col] == '#' {
		if direction == 3 {
			direction = 0
		} else {
			direction++
		}
		return move(matrix, sRow, sCol, dr, dc, direction)
	}
	return matrix
}

func trackGuardAndTrap(matrix [][]byte, sRow, sCol int) int {
	var trap int
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}
	direction := 0
	_, trap = moveAndTrap(matrix, sRow, sCol, dr, dc, direction, trap)
	return trap
}

func moveAndTrap(matrix [][]byte, sRow, sCol int, dr, dc []int, direction, trap int) ([][]byte, int) {
	row := sRow + dr[direction]
	col := sCol + dc[direction]
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		trap++
		return matrix, trap
	}
	if matrix[row][col] == '.' || matrix[row][col] == 'X' {
		var temp int
		if direction == 3 {
			temp = 0
		} else {
			temp = direction + 1
		}
		if matrix[row+dr[temp]][col+dc[temp]] == 'X' {
			trap++
		}
		matrix[row][col] = 'X'
		return moveAndTrap(matrix, row, col, dr, dc, direction, trap)
	}
	if matrix[row][col] == '#' {
		if direction == 3 {
			direction = 0
		} else {
			direction++
		}
		return moveAndTrap(matrix, sRow, sCol, dr, dc, direction, trap)
	}
	return matrix, trap
}
