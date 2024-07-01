package X0492_ConstructTheRectangle

import (
	"encoding/json"
	"fmt"
	"math"
)

func Main() {
	area := 4

	result := constructRectangle(area)
	resultString, _ := json.Marshal(result)
	fmt.Println("Result Construct the Rectangle: " + string(resultString))
}

func constructRectangle(area int) []int {
	left := math.Sqrt(float64(area))

	if area == 1 {
		return []int{1, 1}
	}
	if left == float64(int(left)) {
		return []int{int(left), int(left)}
	}

	for i := int(left); i > 0; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}

	return []int{area, 1}
}
