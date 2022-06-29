package X0118_PascalsTriangle

import (
	"encoding/json"
	"fmt"
)

func Main() {
	rows := 5

	output := generate(rows)
	result, _ := json.Marshal(output)
	fmt.Println("Result Pascal's Triangle: " + string(result))

}

func generate(numRows int) [][]int {
	if numRows < 1 {
		return [][]int{}
	}
	result := make([][]int, numRows)
	result[0] = []int{1}

	numFields := 2
	for row := 1; row < numRows; row++ {
		result[row] = make([]int, numFields)
		for field := 0; field < numFields; field++ {
			leftAncestor := 0
			if field-1 >= 0 {
				leftAncestor = result[row-1][field-1]
			}
			rightAncestor := 0
			if field < len(result[row-1]) {
				rightAncestor = result[row-1][field]
			}
			result[row][field] = leftAncestor + rightAncestor
		}
		numFields++
	}
	return result
}
