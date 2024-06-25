package X0476_NumberComplement

import (
	"fmt"
	"math/bits"
	"strconv"
)

func Main() {
	num := 5

	result := findComplement(num)
	fmt.Println("Result Number Complement: " + strconv.Itoa(result))
}

func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	mask := (1 << uint(63-bits.LeadingZeros(uint(num)))) - 1
	return mask &^ num
}

func findComplementAlternative(num int) int {
	if num == 0 {
		return 1
	}
	mask := 1
	for mask <= num {
		mask <<= 1
	}
	return mask - 1 - num
}
