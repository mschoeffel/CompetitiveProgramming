package X0448_FindAllNumbersDisappearedInAnArray

import (
	"encoding/json"
	"fmt"
)

func Main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	result := findDisappearedNumbers(nums)
	output, _ := json.Marshal(result)
	fmt.Println("Result Find All Numbers Disappeared in an Array: " + string(output))

}

func findDisappearedNumbers(nums []int) []int {
	occ := make([]bool, len(nums))

	for _, num := range nums {
		occ[num-1] = true
	}

	result := []int{}
	for i, val := range occ {
		if !val {
			result = append(result, i+1)
		}
	}
	return result
}
