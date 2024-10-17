package X0670_MaximumSwap

import (
	"fmt"
	"strconv"
)

func Main() {
	example1 := 2736

	result := maximumSwap(example1)
	fmt.Println("Result Maximum Swap: " + strconv.Itoa(result))
}

func maximumSwap(num int) int {
	numStr := strconv.Itoa(num)
	numArr := []rune(numStr)
	numLen := len(numArr)

	if numLen <= 1 {
		return num
	}

	lastBiggestIndex := 0
	lastBiggestNum := '/'

	previousNum := numArr[0]
	for i := 1; i < numLen; i++ {
		if numArr[i] > previousNum && numArr[i] >= lastBiggestNum || (numArr[i] == previousNum && numArr[i] == lastBiggestNum) {
			lastBiggestIndex = i
			lastBiggestNum = numArr[i]
		}
		previousNum = numArr[i]
	}

	swapNum := numArr[0]
	swapNumIndex := 0

	for i := 0; i < numLen; i++ {
		if numArr[i] < lastBiggestNum {
			swapNum = numArr[i]
			swapNumIndex = i
			break
		}
	}

	if lastBiggestIndex > swapNumIndex {
		numArr[lastBiggestIndex] = swapNum
		numArr[swapNumIndex] = lastBiggestNum
	} else {
		return num
	}

	result, _ := strconv.Atoi(string(numArr))
	return result
}
