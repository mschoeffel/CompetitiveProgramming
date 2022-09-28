package X0349_IntersectionOfTwoArrays

import (
	"encoding/json"
	"fmt"
)

func Main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	result := intersection(nums1, nums2)
	output, _ := json.Marshal(result)
	fmt.Println("Result Intersection of Two Arrays: " + string(output))

}

func intersection(nums1 []int, nums2 []int) []int {
	var count = map[int]bool{}
	var result []int

	for _, num := range nums1 {
		count[num] = true
	}

	for _, num := range nums2 {
		if count[num] {
			result = append(result, num)
			count[num] = false
		}
	}

	return result
}
