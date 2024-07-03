package X0496_NextGreaterElementI

import (
	"encoding/json"
	"fmt"
)

func Main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	result := nextGreaterElement(nums1, nums2)
	resultString, _ := json.Marshal(result)
	fmt.Println("Result Next Greater Element I: " + string(resultString))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	store := make(map[int]int, len(nums1))
	var stack []int
	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			store[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	for i, num := range nums1 {
		if v, ok := store[num]; ok {
			nums1[i] = v
			continue
		}
		nums1[i] = -1
	}
	return nums1
}
