package X0035_SearchInsertPosition

import (
	"fmt"
	"sort"
	"strconv"
)

func Main() {
	nums := []int{1, 3, 5, 6}
	target := 5

	result := searchInsert(nums, target)
	fmt.Println("Result Search Insert: " + strconv.Itoa(result))
	result = searchInsertBuildIn(nums, target)
	fmt.Println("Result Search Insert Build In: " + strconv.Itoa(result))

}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return start
}

func searchInsertBuildIn(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
}
