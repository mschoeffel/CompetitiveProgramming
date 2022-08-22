package X0228_SummaryRanges

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func Main() {
	nums := []int{0, 1, 2, 4, 5, 7}

	output := summaryRanges(nums)
	result, _ := json.Marshal(output)
	fmt.Println("Result Summary Ranges: " + string(result))

}

func summaryRanges(nums []int) []string {
	if len(nums) <= 0 {
		return []string{}
	}

	var result []string
	var start int
	startBool := false
	var previousNum int
	previousNumBool := false

	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]
		if !startBool {
			start = currentNum
			startBool = true
		}
		if previousNumBool && previousNum+1 != currentNum {
			if start == previousNum {
				result = append(result, strconv.Itoa(previousNum))
			} else {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(previousNum))
			}
			previousNumBool = false
			start = currentNum
		}
		previousNum = currentNum
		previousNumBool = true
	}
	if start == previousNum {
		result = append(result, strconv.Itoa(previousNum))
	} else {
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(previousNum))
	}
	return result
}
