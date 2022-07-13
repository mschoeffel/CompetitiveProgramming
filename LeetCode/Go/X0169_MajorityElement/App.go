package X0169_MajorityElement

import (
	"fmt"
	"sort"
	"strconv"
)

func Main() {
	nums := []int{3, 2, 3}

	result := majorityElement(nums)
	fmt.Println("Result Majority Element: " + strconv.Itoa(result))

	result = majorityElementSort(nums)
	fmt.Println("Result Majority Element Sort: " + strconv.Itoa(result))
}

func majorityElement(nums []int) int {
	count := 0
	var maxNum int
	for _, num := range nums {
		if count == 0 {
			maxNum = num
		}
		if num == maxNum {
			count++
		} else {
			count--
		}
	}
	return maxNum
}

func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	maxCount := 0
	maxNum := 0
	currCount := 0
	currNum := 0
	for _, x := range nums {
		if x == currNum {
			currCount++
		} else {
			currNum = x
			currCount = 1
		}
		if currCount > maxCount {
			maxCount = currCount
			maxNum = currNum
		}
	}
	return maxNum
}
