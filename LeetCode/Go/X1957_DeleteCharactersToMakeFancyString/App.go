package X1957_DeleteCharactersToMakeFancyString

import (
	"fmt"
)

func Main() {
	s := "leeetcode"

	result := makeFancyString(s)
	fmt.Println("Result Delete Characters to Make Fancy String: " + result)
}

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}

	runeArray := []rune(s)

	j := 2

	for i := 2; i < len(runeArray); i++ {
		if runeArray[i] != runeArray[j-1] || runeArray[i] != runeArray[j-2] {
			runeArray[j] = runeArray[i]
			j++
		}
	}

	return string(runeArray[:j])
}
