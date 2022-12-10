package Day10

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"regexp"
	"strconv"
)

func Main() {
	program := Utils.ReadFileByLinesString("\\Day10\\program.txt")
	fmt.Println("Day 10: Result part one: " + strconv.Itoa(Day10Part1(program)))
	fmt.Println("Day 10: Result part two: ")
	Day10Part2(program)
}

func Day10Part1(program []string) int {
	regCommand := regexp.MustCompile(`[a-zA-z]+`)
	regNum := regexp.MustCompile(`[-\d]+`)

	x := 1
	cycle := 0
	save := []int{}

	for _, programLine := range program {
		command := regCommand.FindString(programLine)
		value, _ := strconv.Atoi(regNum.FindString(programLine))

		switch command {
		case "addx":
			for c := 0; c < 2; c++ {
				cycle++
				if cycle == 20 || (cycle-20)%40 == 0 {
					save = append(save, x)
				}
			}
			x += value
			break
		case "noop":
			cycle++
			if cycle == 20 || (cycle-20)%40 == 0 {
				save = append(save, x)
			}
			break
		}
	}

	sum := 0
	for i, v := range save {
		sum += (i*40 + 20) * v
	}
	return sum
}

func Day10Part2(program []string) {
	regCommand := regexp.MustCompile(`[a-zA-z]+`)
	regNum := regexp.MustCompile(`[-\d]+`)
	posX := 0
	posY := 0
	screen := [6][40]rune{}

	x := 1
	cycle := 0

	for _, programLine := range program {
		command := regCommand.FindString(programLine)
		value, _ := strconv.Atoi(regNum.FindString(programLine))

		switch command {
		case "addx":
			for c := 0; c < 2; c++ {
				if posX == x || posX+1 == x || posX-1 == x {
					screen[posY][posX] = '#'
				} else {
					screen[posY][posX] = '.'
				}
				cycle++
				posX++
				if cycle%40 == 0 {
					posY++
					posX = 0
				}
			}
			x += value
			break
		case "noop":
			if posX == x || posX+1 == x || posX-1 == x {
				screen[posY][posX] = '#'
			} else {
				screen[posY][posX] = '.'
			}
			cycle++
			posX++
			if cycle%40 == 0 {
				posY++
				posX = 0
			}

			break
		}
	}

	for sY := 0; sY < len(screen); sY++ {
		for sX := 0; sX < len(screen[0]); sX++ {
			print(string(screen[sY][sX]))
		}
		println()
	}

}
