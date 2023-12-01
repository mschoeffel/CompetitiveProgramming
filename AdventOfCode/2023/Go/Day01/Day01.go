package Day01

import (
	"AdventOfCode-2023/Utils"
	"fmt"
	"strconv"
	"unicode"
)

func Main() {
	document := Utils.ReadFileByLinesString("\\Day01\\document.txt")
	fmt.Println("Day 1: Result part one: " + strconv.Itoa(Day1Part1(document)))
}

func Day1Part1(document []string) int {
	sum := 0
	for _, documentLine := range document {
		firstDigit := -1
		lastDigit := -1

		for _, runeValue := range documentLine {
			if unicode.IsDigit(runeValue) {
				firstDigit = int(runeValue - '0')
				break
			}
		}
		for i := len(documentLine) - 1; i >= 0; i-- {
			runeValue := rune(documentLine[i])
			if unicode.IsDigit(runeValue) {
				lastDigit = int(runeValue - '0')
				break
			}
		}

		if lastDigit != -1 && firstDigit != -1 {
			sum += firstDigit*10 + lastDigit
		} else if firstDigit != -1 {
			sum += firstDigit
		}
	}
	return sum
}
