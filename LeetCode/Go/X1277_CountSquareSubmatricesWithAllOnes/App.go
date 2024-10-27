package X1277_CountSquareSubmatricesWithAllOnes

import (
	"fmt"
	"strconv"
)

func Main() {
	matrix := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}

	result := countSquares(matrix)
	fmt.Println("Result Count Square Submatrices with All Ones: " + strconv.Itoa(result))
}

func countSquares(matrix [][]int) int {
	numberRows := len(matrix)
	numberCols := len(matrix[0])

	dp := make([][]int, numberRows)
	for i := range dp {
		dp[i] = make([]int, numberCols)
	}

	totalCountSquares := 0

	for i := 0; i < numberRows; i++ {
		dp[i][0] = matrix[i][0]
		totalCountSquares += dp[i][0]
	}

	for j := 1; j < numberCols; j++ {
		dp[0][j] = matrix[0][j]
		totalCountSquares += dp[0][j]
	}

	for i := 1; i < numberRows; i++ {
		for j := 1; j < numberCols; j++ {
			if matrix[i][j] == 1 {
				localMin := dp[i][j-1]
				if dp[i-1][j] < localMin {
					localMin = dp[i-1][j]
				}
				if dp[i-1][j-1] < localMin {
					localMin = dp[i-1][j-1]
				}
				dp[i][j] = 1 + localMin
			}
			totalCountSquares += dp[i][j]
		}
	}

	return totalCountSquares
}
