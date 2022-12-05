package Day04

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"strconv"
	"strings"
)

func Main() {
	sections := Utils.ReadFileByLinesString("\\Day04\\sections.txt")
	fmt.Println("Day 4: Result part one: " + strconv.Itoa(Day4Part1(sections)))
	fmt.Println("Day 4: Result part two: " + strconv.Itoa(Day4Part2(sections)))
}

func Day4Part1(sections []string) int {
	count := 0

	for _, section := range sections {
		s := strings.Split(section, ",")
		s1 := strings.Split(s[0], "-")
		s2 := strings.Split(s[1], "-")

		firstStart, _ := strconv.Atoi(s1[0])
		firstEnd, _ := strconv.Atoi(s1[1])
		secondStart, _ := strconv.Atoi(s2[0])
		secondEnd, _ := strconv.Atoi(s2[1])

		if (firstStart >= secondStart && firstEnd <= secondEnd) || (secondStart >= firstStart && secondEnd <= firstEnd) {
			count++
		}
	}

	return count
}

func Day4Part2(sections []string) int {
	count := 0

	for _, section := range sections {
		s := strings.Split(section, ",")
		s1 := strings.Split(s[0], "-")
		s2 := strings.Split(s[1], "-")

		firstStart, _ := strconv.Atoi(s1[0])
		firstEnd, _ := strconv.Atoi(s1[1])
		secondStart, _ := strconv.Atoi(s2[0])
		secondEnd, _ := strconv.Atoi(s2[1])

		if (firstStart >= secondStart && firstStart <= secondEnd) || (secondStart >= firstStart && secondStart <= firstEnd) {
			count++
		}
	}

	return count
}
