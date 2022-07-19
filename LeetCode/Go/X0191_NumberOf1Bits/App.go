package X0191_NumberOf1Bits

import (
	"fmt"
	"strconv"
)

func Main() {
	var num uint32 = 43261596

	result := hammingWeight(num)
	fmt.Println("Result Number of 1 Bits: " + strconv.Itoa(result))
}

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i <= 31; i++ {
		if num&(1<<i) > 0 {
			count++
		}
	}
	return count
}
