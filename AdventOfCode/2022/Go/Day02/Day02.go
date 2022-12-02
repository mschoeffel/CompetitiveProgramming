package Day02

import (
	"AdventOfCode-2022/Utils"
	"fmt"
	"strconv"
)

func Main() {
	inventory := Utils.ReadFileByLinesString("\\Day02\\guide.txt")
	fmt.Println("Day 2: Result part one: " + strconv.Itoa(Day2Part1(inventory)))
	fmt.Println("Day 2: Result part two: " + strconv.Itoa(Day2Part2(inventory)))
}

func Day2Part1(inventory []string) int {
	score := 0
	for i := 0; i < len(inventory); i++ {
		opponentChar := rune(inventory[i][0])
		// A = Rock
		// B = Paper
		// C = Scissors
		ownChar := rune(inventory[i][2])
		// X = Rock = 1
		// Y = Paper = 2
		// Z = Scissors = 3
		// Win = 6
		// Draw = 3
		// Loss = 0

		switch opponentChar {
		case 'A':
			switch ownChar {
			case 'X':
				score += 1 + 3
				break
			case 'Y':
				score += 2 + 6
				break
			case 'Z':
				score += 3
				break
			}
			break
		case 'B':
			switch ownChar {
			case 'X':
				score += 1
				break
			case 'Y':
				score += 2 + 3
				break
			case 'Z':
				score += 3 + 6
				break
			}
			break
		case 'C':
			switch ownChar {
			case 'X':
				score += 1 + 6
				break
			case 'Y':
				score += 2
				break
			case 'Z':
				score += 3 + 3
				break
			}
			break
		}
	}
	return score
}

func Day2Part2(inventory []string) int {
	score := 0
	for i := 0; i < len(inventory); i++ {
		opponentChar := rune(inventory[i][0])
		// A = Rock = 1
		// B = Paper = 2
		// C = Scissors = 3
		ownChar := rune(inventory[i][2])
		// X = Loose
		// Y = Draw
		// Z = Win
		// Win = 6
		// Draw = 3
		// Loss = 0

		switch opponentChar {
		case 'A':
			switch ownChar {
			case 'X':
				score += 3 + 0
				break
			case 'Y':
				score += 1 + 3
				break
			case 'Z':
				score += 2 + 6
				break
			}
			break
		case 'B':
			switch ownChar {
			case 'X':
				score += 1 + 0
				break
			case 'Y':
				score += 2 + 3
				break
			case 'Z':
				score += 3 + 6
				break
			}
			break
		case 'C':
			switch ownChar {
			case 'X':
				score += 2 + 0
				break
			case 'Y':
				score += 3 + 3
				break
			case 'Z':
				score += 1 + 6
				break
			}
			break
		}
	}
	return score
}
