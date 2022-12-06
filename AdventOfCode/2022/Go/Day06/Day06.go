package Day06

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"strconv"
)

func Main() {
	datastream := Utils.ReadFileByLinesString("\\Day06\\datastream.txt")[0]
	fmt.Println("Day 6: Result part one: " + strconv.Itoa(Day6Part1(datastream)))
	fmt.Println("Day 6: Result part two: " + strconv.Itoa(Day6Part2(datastream)))
}

func Day6Part1(datastream string) int {

	for i, v := range datastream {
		if i < 3 {
			continue
		}
		occur := make(map[rune]int)
		occur[v]++
		occur[rune(datastream[i-1])]++
		occur[rune(datastream[i-2])]++
		occur[rune(datastream[i-3])]++
		single := true
		for _, o := range occur {
			if o > 1 {
				single = false
				break
			}
		}
		if single {
			return i + 1
		}
	}

	return 0
}

func Day6Part2(datastream string) int {

	for i, _ := range datastream {
		if i < 14 {
			continue
		}
		occur := make(map[rune]int)
		for j := 0; j < 14; j++ {
			occur[rune(datastream[i-j])]++
		}
		single := true
		for _, o := range occur {
			if o > 1 {
				single = false
				break
			}
		}
		if single {
			return i + 1
		}
	}

	return 0
}
