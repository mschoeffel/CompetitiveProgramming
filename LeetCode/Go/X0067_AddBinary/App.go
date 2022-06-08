package X0067_AddBinary

import (
	"fmt"
)

func Main() {
	a := "11011"
	b := "11"

	result := addBinary(a, b)
	fmt.Println("Result Add Binary: " + result)

}

func addBinary(a string, b string) string {
	pointerA := len(a) - 1
	pointerB := len(b) - 1
	overflow := 0
	result := ""

	for pointerA >= -1 || pointerB >= -1 {
		var numA rune
		if pointerA >= 0 {
			numA = rune(a[pointerA])
		} else {
			numA = '0'
		}

		var numB rune
		if pointerB >= 0 {
			numB = rune(b[pointerB])
		} else {
			numB = '0'
		}

		if numA == '1' && numB == '1' {
			if overflow > 0 {
				overflow = 0
				result = "1" + result
			} else {
				result = "0" + result
			}
			overflow += 1

		} else if numA == '0' && numB == '0' {
			if overflow > 0 {
				overflow = 0
				result = "1" + result
			} else {
				result = "0" + result
			}
		} else {
			if overflow > 0 {
				overflow = 1
				result = "0" + result
			} else {
				result = "1" + result
			}
		}
		pointerA--
		pointerB--
	}
	if result[0] == '0' {
		return result[1:]
	}
	return result
}
