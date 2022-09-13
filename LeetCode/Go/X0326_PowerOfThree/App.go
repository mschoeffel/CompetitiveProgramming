package X0326_PowerOfThree

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 4

	output := isPowerOfThree(num)
	fmt.Println("Result Power of Three: " + strconv.FormatBool(output))

}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || (n%3) != 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}
