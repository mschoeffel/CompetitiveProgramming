package X0374_GuessNumberHigherOrLower

import (
	"fmt"
	"math"
	"strconv"
)

func Main() {
	n := 10

	result := guessNumber(n)
	fmt.Println("Result Guess Number Higher or Lower: " + strconv.Itoa(result))

}

func guessNumber(n int) int {
	lower := 0
	upper := n
	nextGuess := int(math.Ceil(float64(n) / 2))
	answer := guess(nextGuess)
	for answer != 0 {
		if answer == -1 {
			upper = nextGuess
		} else if answer == 1 {
			lower = nextGuess
		}
		nextGuess = int(math.Ceil(float64(upper-lower)/2)) + lower
		answer = guess(nextGuess)
	}
	return nextGuess
}

func guess(num int) int {
	pick := 6
	if num > pick {
		return -1
	} else if num < pick {
		return 1
	} else {
		return 0
	}
}
