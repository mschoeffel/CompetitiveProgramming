package X0461_HammingDistance

import (
	"fmt"
	"math/bits"
	"strconv"
)

func Main() {
	x := 1
	y := 4

	result := hammingDistance(x, y)
	fmt.Println("Result Hamming Distance: " + strconv.Itoa(result))
}

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
