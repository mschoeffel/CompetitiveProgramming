package X0088_MergeSortedArray

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
)

func Main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	nums1Temp := make([]int, len(nums1))
	copy(nums1Temp, nums1)
	merge(nums1Temp, m, nums2, n)
	result, _ := json.Marshal(nums1Temp)
	fmt.Println("Result Merge Two Sorted Lists: " + string(result))

	mergeOptimized(nums1Temp, m, nums2, n)
	result, _ = json.Marshal(nums1Temp)
	fmt.Println("Result Merge Two Sorted Lists Optimized: " + string(result))

	mergeBuiltIn(nums1Temp, m, nums2, n)
	result, _ = json.Marshal(nums1Temp)
	fmt.Println("Result Merge Two Sorted Lists BuiltIn: " + string(result))

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	pointerNums1 := m - 1
	pointerNums2 := n - 1
	for i := m + n - 1; i >= 0; i-- {
		currentMaxNums1 := -math.MaxInt32
		if pointerNums1 >= 0 {
			currentMaxNums1 = nums1[pointerNums1]
		}

		currentMaxNums2 := -math.MaxInt32
		if pointerNums2 >= 0 {
			currentMaxNums2 = nums2[pointerNums2]
		}

		if currentMaxNums1 >= currentMaxNums2 {
			nums1[i] = currentMaxNums1
			pointerNums1--
		} else {
			nums1[i] = currentMaxNums2
			pointerNums2--
		}
	}
}

func mergeOptimized(nums1 []int, m int, nums2 []int, n int) {

	pointerNums1 := m - 1
	pointerNums2 := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if pointerNums1 < 0 {
			nums1[i] = nums2[pointerNums2]
			pointerNums2--
			continue
		}

		if pointerNums2 < 0 {
			continue
		}

		if nums1[pointerNums1] >= nums2[pointerNums2] {
			nums1[i] = nums1[pointerNums1]
			pointerNums1--
		} else {
			nums1[i] = nums2[pointerNums2]
			pointerNums2--
		}
	}
}

func mergeBuiltIn(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
