package X2684_MaximumNumberOfMovesInAGrid

import (
	"fmt"
	"strconv"
)

func Main() {
	grid := [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}

	result := maxMoves(grid)
	fmt.Println("Result Maximum Number of Moves in a Grid: " + strconv.Itoa(result))
}

func maxMoves(grid [][]int) int {
	numRows, numCols := len(grid), len(grid[0])
	dp := make([][]int, numRows)
	for i := range dp {
		dp[i] = make([]int, numCols)
	}
	maxPossibleMoves := 0

	for col := numCols - 2; col >= 0; col-- {
		for row := 0; row < numRows; row++ {
			if row > 0 && grid[row][col] < grid[row-1][col+1] {
				dp[row][col] = max(dp[row][col], dp[row-1][col+1]+1)
			}
			if grid[row][col] < grid[row][col+1] {
				dp[row][col] = max(dp[row][col], dp[row][col+1]+1)
			}
			if row < numRows-1 && grid[row][col] < grid[row+1][col+1] {
				dp[row][col] = max(dp[row][col], dp[row+1][col+1]+1)
			}
		}
	}

	for row := 0; row < numRows; row++ {
		maxPossibleMoves = max(maxPossibleMoves, dp[row][0])
	}
	return maxPossibleMoves
}
