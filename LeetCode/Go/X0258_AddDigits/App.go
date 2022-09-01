package X0258_AddDigits

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 38

	output := addDigits(num)
	fmt.Println("Result Add Digits: " + strconv.Itoa(output))

	output = addDigitsMath(num)
	fmt.Println("Result Add Digits Math: " + strconv.Itoa(output))

}

func addDigits(num int) int {
	sum := 10
	for sum >= 10 {
		sum = 0
		digits := strconv.Itoa(num)
		for i := 0; i < len(digits); i++ {
			sum += int(digits[i] - '0')
		}
		num = sum
	}
	return sum
}

func addDigitsMath(num int) int {
	return 1 + (num-1)%9
}
