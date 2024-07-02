package X0495_TeemoAttacking

import (
	"fmt"
	"strconv"
)

func Main() {
	timeSeries := []int{1, 4}
	duration := 2

	result := findPoisonedDuration(timeSeries, duration)
	fmt.Println("Result Teemo Attacking: " + strconv.Itoa(result))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	pool := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		pool += min(timeSeries[i+1]-timeSeries[i], duration)
	}
	return pool + duration
}
