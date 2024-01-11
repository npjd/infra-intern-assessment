package main


func printGrid(grid [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			print(grid[i][j], " ")
		}
		println()
	}

}

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


func isValid(grid [][]int, x int, y int, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[x][i] == num {
			return false
		}
		if grid[i][y] == num {
			return false
		}
	}
	
	for i := x/3; i < x/3+3; i++ {
		for j := y/3; j < y/3+3; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}

	return true
}




func SolveSudoku(grid [][]int) [][]int {
	
	var x, y int
	
	x, y = findNextEmpty(grid)

	if x == -1 && y == -1 {
		return grid
	}

	for num := 1; num <= 9; num++ {
		if isValid(grid, x, y, num)  {
			grid[x][y] = num
			if SolveSudoku(grid) != nil {
				return grid
			}

			grid[x][y] = 0
			
		}
	}

	return nil

}




