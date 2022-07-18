package X0190_ReverseBits

import (
	"fmt"
	"strconv"
)

func Main() {
	var num uint32 = 43261596

	result := reverseBits(num)
	fmt.Println("Result Reverse Bits: " + strconv.FormatUint(uint64(result), 10))
}

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		bi := uint32(1) << uint(i)
		bj := uint32(1) << uint(32-i-1)
		if num&bi != 0 {
			res |= bj
		}
	}
	return res
}
