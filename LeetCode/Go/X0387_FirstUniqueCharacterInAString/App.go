package X0387_FirstUniqueCharacterInAString

import (
	"fmt"
	"strconv"
)

func Main() {
	s := "leetcode"

	result := firstUniqChar(s)
	fmt.Println("Result First Unique Character in a String: " + strconv.Itoa(result))

}

func firstUniqChar(s string) int {
	occur := make(map[rune]int)

	for i, v := range s {
		if _, ok := occur[v]; !ok {
			occur[v] = i
		} else {
			occur[v] = -1
		}
	}

	min := -1
	for _, v := range occur {
		if min == -1 || (v != -1 && v < min) {
			min = v
		}
	}

	return min
}
