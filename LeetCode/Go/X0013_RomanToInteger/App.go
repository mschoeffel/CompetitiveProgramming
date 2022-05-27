package X0013_RomanToInteger

import (
	"fmt"
	"strconv"
)

func Main() {
	roman := "MCMXCIV"

	result := romanToInt(roman)
	fmt.Println("Result Roman Integer: " + strconv.Itoa(result))

}

func romanToInt(s string) int {
	value := 0
	for i := 0; i < len(s); i++ {
		var followingChar byte
		if i+1 < len(s) {
			followingChar = s[i+1]
		}
		switch s[i] {
		case 'M':
			value += 1000
			break
		case 'D':
			value += 500
			break
		case 'C':
			if followingChar == 'D' {
				value += 400
				i++
			} else if followingChar == 'M' {
				value += 900
				i++
			} else {
				value += 100
			}
			break
		case 'L':
			value += 50
			break
		case 'X':
			if followingChar == 'L' {
				value += 40
				i++
			} else if followingChar == 'C' {
				value += 90
				i++
			} else {
				value += 10
			}
			break
		case 'V':
			value += 5
			break
		case 'I':
			if followingChar == 'V' {
				value += 4
				i++
			} else if followingChar == 'X' {
				value += 9
				i++
			} else {
				value += 1
			}
			break
		}
	}
	return value
}
