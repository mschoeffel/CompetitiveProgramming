package X0017_LetterCombinationsOfAPhoneNumber

import (
	"encoding/json"
	"fmt"
)

func Main() {
	example1 := "23"

	result, _ := json.Marshal(letterCombinations(example1))
	fmt.Println("Result Letter Combinations of a Phone Number: " + string(result))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digitMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	result := []string{""}
	for _, digit := range digits {
		var temp []string
		for _, letter := range digitMap[string(digit)] {
			for _, res := range result {
				temp = append(temp, res+string(letter))
			}
		}
		result = temp
	}
	return result
}
