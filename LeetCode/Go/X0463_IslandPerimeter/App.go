package X0463_IslandPerimeter

import (
	"fmt"
	"strconv"
)

func Main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}

	result := islandPerimeter(grid)
	fmt.Println("Result Hamming Distance: " + strconv.Itoa(result))
}

func islandPerimeter(grid [][]int) int {
	perimeter := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			cell := grid[i][j]
			if cell == 1 {
				if i == 0 {
					perimeter += 1
				}
				if j == 0 {
					perimeter += 1
				}
				if i == len(grid)-1 {
					perimeter += 1
				}
				if j == len(grid[i])-1 {
					perimeter += 1
				}

				if i > 0 && grid[i-1][j] == 0 {
					perimeter += 1
				}

				if j < len(grid[i])-1 && grid[i][j+1] == 0 {
					perimeter += 1
				}

				if i < len(grid)-1 && grid[i+1][j] == 0 {
					perimeter += 1
				}

				if j > 0 && grid[i][j-1] == 0 {
					perimeter += 1
				}
			}
		}
	}

	return perimeter
}
