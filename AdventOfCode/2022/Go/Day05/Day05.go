package Day05

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"regexp"
	"strconv"
)

func Main() {
	crane := Utils.ReadFileByLinesString("\\Day05\\crane.txt")
	fmt.Println("Day 5: Result part one: " + Day5Part1(crane))
	fmt.Println("Day 5: Result part two: " + Day5Part2(crane))
}

func Day5Part1(sections []string) string {
	const numStacks = 9
	const highestStack = 8

	stacks := [numStacks][]rune{}
	for stack := 0; stack < numStacks; stack++ {
		for line := highestStack - 1; line >= 0; line-- {
			idx := 1 + (stack * 4)
			if idx < len(sections[line]) {
				r := rune(sections[line][idx])
				if r != ' ' {
					stacks[stack] = append(stacks[stack], r)
				}
			}
		}
	}

	reg := regexp.MustCompile(`\d+`)

	for line := highestStack + 2; line < len(sections); line++ {
		info := reg.FindAllString(sections[line], 3)
		amount, _ := strconv.Atoi(info[0])
		source, _ := strconv.Atoi(info[1])
		destination, _ := strconv.Atoi(info[2])
		for i := 0; i < amount; i++ {
			item := stacks[source-1][len(stacks[source-1])-1]
			stacks[destination-1] = append(stacks[destination-1], item)
			stacks[source-1] = stacks[source-1][:len(stacks[source-1])-1]
		}
	}

	top := ""
	for _, stack := range stacks {
		top += string(stack[len(stack)-1])
	}
	return top
}

func Day5Part2(sections []string) string {
	const numStacks = 9
	const highestStack = 8

	stacks := [numStacks][]rune{}
	for stack := 0; stack < numStacks; stack++ {
		for line := highestStack - 1; line >= 0; line-- {
			idx := 1 + (stack * 4)
			if idx < len(sections[line]) {
				r := rune(sections[line][idx])
				if r != ' ' {
					stacks[stack] = append(stacks[stack], r)
				}
			}
		}
	}

	reg := regexp.MustCompile(`\d+`)

	for line := highestStack + 2; line < len(sections); line++ {
		info := reg.FindAllString(sections[line], 3)
		amount, _ := strconv.Atoi(info[0])
		source, _ := strconv.Atoi(info[1])
		destination, _ := strconv.Atoi(info[2])

		item := stacks[source-1][len(stacks[source-1])-amount:]
		stacks[destination-1] = append(stacks[destination-1], item...)
		stacks[source-1] = stacks[source-1][:len(stacks[source-1])-amount]

	}

	top := ""
	for _, stack := range stacks {
		top += string(stack[len(stack)-1])
	}
	return top
}
