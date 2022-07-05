package X0136_SingleNumber

import (
	"fmt"
	"strconv"
)

func Main() {
	input := []int{2, 2, 1}

	result := singleNumber(input)
	fmt.Println("Result Single Number: " + strconv.Itoa(result))
	fmt.Println("Result Single Number XOr: " + strconv.Itoa(result))

}

func singleNumber(nums []int) int {
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for key, value := range m {
		if value == 1 {
			return key
		}
	}
	return 0
}

func singleNumberXOr(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}
