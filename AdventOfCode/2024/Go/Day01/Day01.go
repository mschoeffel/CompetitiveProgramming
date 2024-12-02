package Day01

import (
	"AdventOfCode-2024/Utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Main() {
	input := Utils.ReadFileByLinesString("\\Day01\\input.txt")
	fmt.Println("Day 1: Result part one: " + strconv.Itoa(Day1Part1(input)))
	fmt.Println("Day 1: Result part two: " + strconv.Itoa(Day1Part2(input)))
}

func Day1Part1(input []string) int {
	leftList := make([]int, len(input))
	rightList := make([]int, len(input))

	for index, element := range input {
		numbers := strings.Split(element, "   ")
		leftList[index], _ = strconv.Atoi(numbers[0])
		rightList[index], _ = strconv.Atoi(numbers[1])
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	distance := 0

	for index := 0; index < len(leftList); index++ {
		absoluteDistance := rightList[index] - leftList[index]
		if absoluteDistance < 0 {
			absoluteDistance = -absoluteDistance
		}
		distance += absoluteDistance
	}

	return distance
}

func Day1Part2(input []string) int {
	leftList := make([]int, len(input))
	rightMap := make(map[int]int)

	for index, element := range input {
		numbers := strings.Split(element, "   ")
		leftList[index], _ = strconv.Atoi(numbers[0])
		rightValue, _ := strconv.Atoi(numbers[1])
		if _, exists := rightMap[rightValue]; exists {
			rightMap[rightValue]++
		} else {
			rightMap[rightValue] = 1
		}
	}

	sum := 0

	for index := 0; index < len(leftList); index++ {
		leftValue := leftList[index]
		if _, exists := rightMap[leftValue]; exists {
			sum += leftValue * rightMap[leftValue]
		}
	}

	return sum
}
