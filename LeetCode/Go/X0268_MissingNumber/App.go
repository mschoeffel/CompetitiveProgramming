package X0268_MissingNumber

import (
	"fmt"
	"strconv"
)

func Main() {
	num := []int{3, 0, 1}

	output := missingNumber(num)
	fmt.Println("Result Missing Number: " + strconv.Itoa(output))

}

func missingNumber(nums []int) int {
	sum := 0
	sumTarget := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sumTarget += i + 1
	}
	return sumTarget - sum
}
