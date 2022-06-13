package X0070_ClimbingStairs

import (
	"fmt"
	"strconv"
)

func Main() {
	stairs := 3

	result := climbStairs(stairs)
	fmt.Println("Result Climb Stairs: " + strconv.Itoa(result))

}

func climbStairs(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
