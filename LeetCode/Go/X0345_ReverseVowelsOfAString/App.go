package X0345_ReverseVowelsOfAString

import (
	"fmt"
)

func Main() {
	s := "hello"

	output := reverseVowels(s)
	fmt.Println("Result Reverse Vowels of a String: " + output)

}

func reverseVowels(s string) string {
	r := []rune(s)
	idx := 0
	switchX := false
	idy := len(r) - 1
	switchY := false
	for idx < idy {
		runeAtX := r[idx]
		runeAtY := r[idy]
		if runeAtX == 'a' || runeAtX == 'e' || runeAtX == 'i' || runeAtX == 'o' || runeAtX == 'u' ||
			runeAtX == 'A' || runeAtX == 'E' || runeAtX == 'I' || runeAtX == 'O' || runeAtX == 'U' {
			switchX = true
		} else {
			idx++
		}

		if runeAtY == 'a' || runeAtY == 'e' || runeAtY == 'i' || runeAtY == 'o' || runeAtY == 'u' ||
			runeAtY == 'A' || runeAtY == 'E' || runeAtY == 'I' || runeAtY == 'O' || runeAtY == 'U' {
			switchY = true
		} else {
			idy--
		}

		if switchX && switchY {
			save := r[idx]
			r[idx] = r[idy]
			r[idy] = save
			switchX = false
			switchY = false
			idx++
			idy--
		}
	}
	return string(r)
}
