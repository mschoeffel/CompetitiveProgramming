package Day03

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"strconv"
)

func Main() {
	contents := Utils.ReadFileByLinesString("\\Day03\\contents.txt")
	fmt.Println("Day 3: Result part one: " + strconv.Itoa(Day3Part1(contents)))
	fmt.Println("Day 3: Result part two: " + strconv.Itoa(Day3Part2(contents)))
}

func Day3Part1(contents []string) int {
	sum := 0

	for i := 0; i < len(contents); i++ {
		compartmentOne := contents[i][:(len(contents[i]) / 2)]
		compartmentTwo := contents[i][(len(contents[i]) / 2):]
		found := false
		for _, r1 := range compartmentOne {
			for _, r2 := range compartmentTwo {
				if r1 == r2 {
					if r1 <= 'Z' {
						sum += int(r1) - 38
						//27-52
						//65-90
					} else {
						sum += int(r1) - 96
						//1-26
						//97-122
					}
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	return sum
}

func Day3Part2(contents []string) int {
	sum := 0

	for i := 0; i < len(contents); i += 3 {
		occur := make(map[rune]int)

		for j := 0; j < 3; j++ {
			backpack := contents[i+j]
			temp := make(map[rune]bool)
			for _, r := range backpack {
				if !temp[r] {
					occur[r]++
					temp[r] = true
				}
			}
		}

		for k, v := range occur {
			if v == 3 {
				if k <= 'Z' {
					sum += int(k) - 38
					//27-52
					//65-90
				} else {
					sum += int(k) - 96
					//1-26
					//97-122
				}
			}
		}
	}
	return sum
}
