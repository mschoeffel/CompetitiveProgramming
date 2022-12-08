package Day08

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"strconv"
)

func Main() {
	input := Utils.ReadFileByLinesString("\\Day08\\map.txt")
	fmt.Println("Day 8: Result part one: " + strconv.Itoa(Day8Part1(input)))
	fmt.Println("Day 8: Result part two: " + strconv.Itoa(Day8Part2(input)))
}

func Day8Part1(sections []string) int {
	visibleTrees := make(map[string]bool)

	for rowIdx := 1; rowIdx < len(sections)-1; rowIdx++ {
		row := sections[rowIdx]
		rowMax := int(row[0] - '0')
		for start := 1; start < len(row)-1; start++ {
			current := int(row[start] - '0')
			if current > rowMax {
				visibleTrees[strconv.Itoa(rowIdx)+"|"+strconv.Itoa(start)] = true
				rowMax = current
			}

		}
		rowMax = int(row[len(row)-1] - '0')
		for end := len(row) - 2; end >= 1; end-- {
			current := int(row[end] - '0')
			if current > rowMax {
				visibleTrees[strconv.Itoa(rowIdx)+"|"+strconv.Itoa(end)] = true
				rowMax = current
			}
		}
	}

	for col := 1; col < len(sections[0])-1; col++ {
		colMax := int(sections[0][col] - '0')
		for top := 1; top < len(sections)-1; top++ {
			current := int(sections[top][col] - '0')
			if current > colMax {
				visibleTrees[strconv.Itoa(top)+"|"+strconv.Itoa(col)] = true
				colMax = current
			}
		}

		colMax = int(sections[len(sections)-1][col] - '0')
		for bot := len(sections) - 2; bot >= 1; bot-- {
			current := int(sections[bot][col] - '0')
			if current > colMax {
				visibleTrees[strconv.Itoa(bot)+"|"+strconv.Itoa(col)] = true
				colMax = current
			}
		}
	}

	return len(visibleTrees) + len(sections)*2 + (len(sections[0])-2)*2
}

func Day8Part2(sections []string) int {
	maxScore := 0

	for row := 1; row < len(sections)-1; row++ {
		for col := 1; col < len(sections[0])-1; col++ {
			score := calcScore(sections, row, col, int(sections[row][col]-'0'))
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func calcScore(sections []string, row int, col int, value int) int {
	scoreTop := 0
	if row > 0 {
		for top := row - 1; top >= 0; top-- {
			valueTemp := int(sections[top][col] - '0')
			if valueTemp < value {
				scoreTop++
			} else {
				scoreTop++
				break
			}
		}
	}

	scoreBot := 0
	if row < len(sections)-1 {
		for bot := row + 1; bot < len(sections); bot++ {
			valueTemp := int(sections[bot][col] - '0')
			if valueTemp < value {
				scoreBot++
			} else {
				scoreBot++
				break
			}
		}
	}

	scoreLeft := 0
	if col > 0 {
		for left := col - 1; left >= 0; left-- {
			valueTemp := int(sections[row][left] - '0')
			if valueTemp < value {
				scoreLeft++
			} else {
				scoreLeft++
				break
			}
		}
	}

	scoreRight := 0
	if col < len(sections[0])-1 {
		for right := col + 1; right < len(sections[0]); right++ {
			valueTemp := int(sections[row][right] - '0')
			if valueTemp < value {
				scoreRight++
			} else {
				scoreRight++
				break
			}
		}
	}

	return scoreTop * scoreBot * scoreRight * scoreLeft
}
