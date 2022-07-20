package X0202_HappyNumber

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 19

	result := isHappy(num)
	fmt.Println("Result Happy Number: " + strconv.FormatBool(result))
}

func isHappy(n int) bool {
	return calc(n, map[int]bool{n: true}) == 1
}

func calc(n int, already map[int]bool) int {
	current := n
	sum := 0
	for current > 0 {
		sum += (current % 10) * (current % 10)
		current = current / 10
	}
	if sum <= 1 {
		return sum
	} else if already[sum] {
		return 0
	}
	already[sum] = true
	return calc(sum, already)
}
