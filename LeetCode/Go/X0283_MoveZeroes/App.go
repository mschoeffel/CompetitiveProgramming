package X0283_MoveZeroes

import (
	"encoding/json"
	"fmt"
)

func Main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)
	result, _ := json.Marshal(nums)
	fmt.Println("Result Move Zeroes: " + string(result))

}

func moveZeroes(nums []int) {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if idx != i {
				nums[idx] = nums[i]
				nums[i] = 0
			}
			idx++
		}
	}
}
