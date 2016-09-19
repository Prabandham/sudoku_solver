//This is a sudoku solver. Which uses the Backtracking alogorithm.
//https://en.wikipedia.org/wiki/Backtracking
package main

import (
	"fmt"
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

func main() {
	var grid = [N][N]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	if solveSudoku(&grid) {
		printGrid(&grid)
	} else {
		fmt.Println("This shit is not happening !!!!")
	}
}
