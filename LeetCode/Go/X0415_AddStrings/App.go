package X0415_AddStrings

import (
	"fmt"
)

func Main() {
	num1 := ""
	num2 := ""

	result := addStrings(num1, num2)
	fmt.Println("Result Add Strings: " + result)

}

func addStrings(num1 string, num2 string) string {
	var result []byte
	carry := 0
	for i := 0; i < len(num1) || i < len(num2); i++ {
		d1, d2 := 0, 0
		if i < len(num1) {
			d1 = int(num1[len(num1)-i-1] - '0')
		}
		if i < len(num2) {
			d2 = int(num2[len(num2)-i-1] - '0')
		}
		d := d1 + d2 + carry
		d, carry = d%10, d/10
		result = append(result, byte(d)+'0')
	}
	if carry > 0 {
		result = append(result, byte(carry)+'0')
	}
	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
