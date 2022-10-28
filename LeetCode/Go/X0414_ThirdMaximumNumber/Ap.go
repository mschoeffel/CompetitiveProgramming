package X0414_ThirdMaximumNumber

import (
	"fmt"
	"math"
	"strconv"
)

func Main() {
	nums := []int{3, 2, 1}

	result := thirdMax(nums)
	fmt.Println("Result Third Maximum Number: " + strconv.Itoa(result))

}

func thirdMax(nums []int) int {
	maxNums := []int{-math.MaxInt, -math.MaxInt, -math.MaxInt}
	for _, num := range nums {
		if num > maxNums[0] {
			maxNums[2] = maxNums[1]
			maxNums[1] = maxNums[0]
			maxNums[0] = num
		} else if num > maxNums[1] && num != maxNums[0] {
			maxNums[2] = maxNums[1]
			maxNums[1] = num
		} else if num > maxNums[2] && num != maxNums[0] && num != maxNums[1] {
			maxNums[2] = num
		}
	}
	if maxNums[2] == -math.MaxInt {
		return maxNums[0]
	}
	return maxNums[2]
}
