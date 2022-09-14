package X0338_CountingBits

import (
	"encoding/json"
	"fmt"
)

func Main() {
	num := 4

	output := countBits(num)
	result, _ := json.Marshal(output)
	fmt.Println("Result Counting Bits: " + string(result))

}

func countBits(n int) []int {
	result := make([]int, n+1)

	for idx := 1; idx <= n; idx++ {
		result[idx] = (idx & 1) + result[idx/2]
	}
	return result
}
