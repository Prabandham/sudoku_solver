//This is a sudoku solver. Which uses the Backtracking alogorithm.
//https://en.wikipedia.org/wiki/Backtracking
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UNASSIGNED int = 0
	N          int = 9
)

func solveSudoku(grid *[N][N]int) bool {
	var row, col int

	var tempGrid [N][N]int

	tempGrid = *grid
	if !findUnassignedLocation(tempGrid, &row, &col) {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isSafe(tempGrid, row, col, num) {
			grid[row][col] = num

			if solveSudoku(grid) {
				return true
			}

			grid[row][col] = UNASSIGNED
		}
	}
	return false
}

func findUnassignedLocation(grid [N][N]int, row *int, col *int) bool {
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			if grid[r][c] == UNASSIGNED {
				*row = r
				*col = c
				return true
			}
		}
	}
	return false
}

func isSafe(grid [N][N]int, row int, col int, num int) bool {
	return !UsedInRow(grid, row, num) &&
		!UsedInColumn(grid, col, num) &&
		!UsedInBox(grid, row-row%3, col-col%3, num)
}

func UsedInRow(grid [N][N]int, row int, num int) bool {
	for col := 0; col < N; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

func UsedInColumn(grid [N][N]int, col int, num int) bool {
	for row := 0; row < N; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

func UsedInBox(grid [N][N]int, boxStartRow int, boxStartCol int, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}

func printGrid(grid *[N][N]int) {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			fmt.Printf("%2d", grid[row][col])
		}
		fmt.Printf("\n")
	}
}

func splitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}
	return subs
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var rawInput string
	for scanner.Scan() {
		rawInput = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	formattedInput := strings.Split(rawInput, " ")
	row := splitSubN(formattedInput[1], 9)
	var grid = [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := range row {
		innerSlice := splitSubN(row[i], 1)
		for j := range innerSlice {
			value, _ := strconv.ParseInt(innerSlice[j], 10, 64)
			grid[i][j] = int(value)
		}
	}

	if solveSudoku(&grid) {
		printGrid(&grid)
	} else {
		fmt.Println("This shit is not happening !!!!")
	}
}

// Run with :-
// echo "rawInput: 000604700706000009000005080070020093800000005430010070050200000300000208002301000 | go run main.go"
