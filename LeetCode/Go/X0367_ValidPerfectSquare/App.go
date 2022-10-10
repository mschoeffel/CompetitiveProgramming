package X0367_ValidPerfectSquare

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 16

	result := isPerfectSquare(num)
	fmt.Println("Result Valid Perfect Square: " + strconv.FormatBool(result))

}

func isPerfectSquare(num int) bool {
	sqrt := float64(num) / 2
	for i := 0; i < 32 && sqrt > 0; i++ {
		sqrtDiff := float64(num) / sqrt
		sqrt = (sqrt + sqrtDiff) / 2
	}
	isQrt := int(sqrt + 0.5)
	return isQrt*isQrt == num
}
