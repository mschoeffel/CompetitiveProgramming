package X0026_RemoveDuplicatesFromSortedArray

import (
	"fmt"
	"strconv"
)

func Main() {
	nums := []int{1, 1, 2}

	result := removeDuplicates(nums)
	fmt.Println("Result Remove Duplicates from sorted Array: " + strconv.Itoa(result))

}

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	pointerWrite := 1
	valueWrite := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != valueWrite {
			nums[pointerWrite] = nums[i]
			pointerWrite++
			valueWrite = nums[i]
		}
	}
	return pointerWrite
}
