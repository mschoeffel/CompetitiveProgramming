package X0350_IntersectionOfTwoArrays2

import (
	"encoding/json"
	"fmt"
)

func Main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	result := intersect(nums1, nums2)
	output, _ := json.Marshal(result)
	fmt.Println("Result Intersection of Two Arrays II: " + string(output))

}

func intersect(nums1 []int, nums2 []int) []int {
	var count = map[int]int{}
	var result []int

	for _, num := range nums1 {
		count[num]++
	}

	for _, num := range nums2 {
		if count[num] > 0 {
			result = append(result, num)
			count[num]--
		}
	}

	return result
}
