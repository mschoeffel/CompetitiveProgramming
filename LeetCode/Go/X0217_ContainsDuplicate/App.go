package X0217_ContainsDuplicate

import (
	"fmt"
	"strconv"
)

func Main() {
	input := []int{1, 2, 3, 1}

	result := containsDuplicate(input)
	fmt.Println("Result Contains Duplicate: " + strconv.FormatBool(result))

}

func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := visited[nums[i]]; ok {
			return true
		}
		visited[nums[i]] = true
	}
	return false
}
