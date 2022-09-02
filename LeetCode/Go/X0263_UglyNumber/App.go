package X0263_UglyNumber

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 6

	output := isUgly(num)
	fmt.Println("Result Ugly Number: " + strconv.FormatBool(output))

}

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}
	return true
}
