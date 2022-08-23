package X0231_PowerOfTwo

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 1

	output := isPowerOfTwo(num)
	fmt.Println("Result Power of Two: " + strconv.FormatBool(output))

}

func isPowerOfTwo(n int) bool {
	return n != 0 && (n&(n-1)) == 0
}
