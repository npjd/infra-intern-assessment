package main

// findNextEmpty function finds the next empty cell (with value 0) in the Sudoku grid.
// It returns the row and column indices of the empty cell.
// If no empty cell is found, it returns (-1, -1).
func findNextEmpty(grid [][]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// isValid function checks if placing a number 'num' at position (x, y) in the Sudoku grid is valid.
// It checks the row, column, and 3x3 subgrid for any conflicts.
// It returns true if the placement is valid, false otherwise.
func isValid(grid [][]int, x, y, num int) bool {

	for i := 0; i < 9; i++ {
		if grid[x][i] == num || grid[i][y] == num {
			return false
		}
	}
	// Check the 3x3 subgrid — integer divison helps us find the starting indices of the subgrid
	for i := 3 * (x / 3); i < 3*(x/3)+3; i++ {
		for j := 3 * (y / 3); j < 3*(y/3)+3; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}

	return true
}

// SolveSudoku function attempts to solve the Sudoku grid using backtracking.
// It returns the solved grid if a solution exists, otherwise it returns nil.
func SolveSudoku(grid [][]int) [][]int {
	var x, y int

	// Find the next empty cell
	x, y = findNextEmpty(grid)

	// If no empty cell is found, the grid is solved
	if x == -1 && y == -1 {
		return grid
	}

	// Try placing numbers 1 to 9 in the empty cell and recursively solve the Sudoku
	for num := 1; num <= 9; num++ {
		if isValid(grid, x, y, num) {
			grid[x][y] = num

			// Recursively attempt to solve the Sudoku
			if SolveSudoku(grid) != nil {
				return grid
			}

			// If the current placement does not lead to a solution, backtrack by resetting the cell value
			grid[x][y] = 0
		}
	}

	// If no valid number can be placed in the current cell, backtrack
	return nil
}
