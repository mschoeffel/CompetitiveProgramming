package X0441_ArrangingCoins

import (
	"fmt"
	"strconv"
)

func Main() {
	n := 5

	result := arrangeCoins(n)
	fmt.Println("Result Arranging Coins: " + strconv.Itoa(result))

}

func arrangeCoins(n int) int {
	idx := 0
	for n >= (idx + 1) {
		idx++
		n -= idx
	}
	return idx
}
