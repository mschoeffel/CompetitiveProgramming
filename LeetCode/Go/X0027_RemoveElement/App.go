package X0027_RemoveElement

import (
	"fmt"
	"strconv"
)

func Main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	result := removeElement(nums, val)
	fmt.Println("Result Remove Element: " + strconv.Itoa(result))

}

func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	pointerWrite := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pointerWrite] = nums[i]
			pointerWrite++
		}
	}
	return pointerWrite
}
