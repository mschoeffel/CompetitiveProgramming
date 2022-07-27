package X0219_ContainsDuplicate2

import (
	"fmt"
	"strconv"
)

func Main() {
	nums := []int{1, 2, 3, 1}
	k := 3

	result := containsNearbyDuplicate(nums, k)
	fmt.Println("Result Contains Duplicate 2: " + strconv.FormatBool(result))

}

func containsNearbyDuplicate(nums []int, k int) bool {
	visited := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if idx, ok := visited[nums[i]]; ok {
			if abs(idx-i) <= k {
				return true
			}
		}
		visited[nums[i]] = i
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
