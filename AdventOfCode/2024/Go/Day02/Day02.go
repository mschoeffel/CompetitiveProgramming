package Day02

import (
	"AdventOfCode-2024/Utils"
	"fmt"
	"strconv"
	"strings"
)

func Main() {
	input := Utils.ReadFileByLinesString("\\Day02\\input.txt")
	fmt.Println("Day 2: Result part one: " + strconv.Itoa(Day2Part1(input)))
}

func Day2Part1(input []string) int {
	data := make([][]int, len(input))

	for index, element := range input {
		numbers := strings.Split(element, " ")
		data[index] = make([]int, len(numbers))
		for i, number := range numbers {
			data[index][i], _ = strconv.Atoi(number)
		}
	}

	unsafeReports := 0

	for _, report := range data {
		isIncreasing := false
		isDecreasing := false

		for i := 1; i < len(report); i++ {
			if report[i] > report[i-1]+3 || report[i] < report[i-1]-3 || report[i] == report[i-1] {
				unsafeReports++
				break
			}
			if report[i] == report[i-1]+1 || report[i] == report[i-1]+2 || report[i] == report[i-1]+3 {
				isIncreasing = true
			}
			if report[i] == report[i-1]-1 || report[i] == report[i-1]-2 || report[i] == report[i-1]-3 {
				isDecreasing = true
			}
			if isIncreasing && isDecreasing {
				unsafeReports++
				break
			}
		}
	}

	return len(data) - unsafeReports
}
