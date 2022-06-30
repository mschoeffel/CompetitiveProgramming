package X0119_PascalsTriangle2

import (
	"encoding/json"
	"fmt"
)

func Main() {
	rows := 3

	output := getRow(rows)
	result, _ := json.Marshal(output)
	fmt.Println("Result Pascal's Triangle II: " + string(result))

}

func getRow(rowIndex int) []int {
	result := make([][]int, rowIndex+1)
	result[0] = []int{1}

	numFields := 2
	for row := 1; row < rowIndex+1; row++ {
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
	return result[rowIndex]
}
