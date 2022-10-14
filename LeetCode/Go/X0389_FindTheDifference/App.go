package X0389_FindTheDifference

import (
	"fmt"
)

func Main() {
	s := "abcd"
	t := "abcde"

	result := findTheDifference(s, t)
	fmt.Println("Result Find the Difference: " + string(result))

}

func findTheDifference(s string, t string) byte {
	occur := make(map[rune]int)

	for _, v := range t {
		occur[v]++
	}

	for _, v := range s {
		occur[v]--
		if occur[v] <= 0 {
			delete(occur, v)
		}
	}

	for k, v := range occur {
		if v > 0 {
			return byte(k)
		}
	}

	return 0
}
