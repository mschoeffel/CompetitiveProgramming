package X2501_LongestSquareStreakInAnArray

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Main() {
	nums := []int{4, 3, 6, 16, 8, 2}

	result := longestSquareStreak(nums)
	fmt.Println("Result Longest Square Streak in an Array: " + strconv.Itoa(result))
}

func longestSquareStreak(nums []int) int {
	mp := make(map[int]int)
	sort.Ints(nums)
	result := -1

	for _, num := range nums {
		sqrt := int(math.Sqrt(float64(num)))

		if sqrt*sqrt != num {
			mp[num] = 1
			continue
		}

		if val, exists := mp[sqrt]; exists {
			mp[num] = val + 1
			if mp[num] > result {
				result = mp[num]
			}
		} else {
			mp[num] = 1
		}
	}

	return result
}
