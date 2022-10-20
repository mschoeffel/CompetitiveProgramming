package X0405_ConvertANumberToHexadecimal

import (
	"fmt"
)

func Main() {
	num := 26

	result := toHex(num)
	fmt.Println("Result Convert a Number to Hexadecimal: " + result)

}

func toHex(num int) string {
	const (
		hexCharacters = "0123456789abcdef"
		mask          = 0xf
	)
	var result = make([]byte, 8)
	var nonzeroIndex = 7

	for i := 7; i >= 0; i-- {
		val := num & mask
		if val > 0 {
			nonzeroIndex = i
		}
		result[i] = hexCharacters[val]
		num >>= 4
	}
	return string(result[nonzeroIndex:])
}
