package X0342_PowerOfFour

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 4

	output := isPowerOfFour(num)
	fmt.Println("Result Power of Four: " + strconv.FormatBool(output))

}

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || (n%4) != 0 {
		return false
	}
	return isPowerOfFour(n / 4)
}

func isPowerOfFourBitPattern(n int) bool {
	const SecondBit int = 0x55555555
	return n > 0 && // Can't have 0 or negatives
		n&(n-1) == 0 && // Checking if only 1 bit is active
		n&SecondBit != 0 // Checking if the bit is in an odd position
}
