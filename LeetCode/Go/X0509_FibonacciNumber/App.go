package X0509_FibonacciNumber

import (
	"fmt"
	"strconv"
)

func Main() {
	example1 := 2

	result := fib(example1)
	fmt.Println("Result Fibonacci Number: " + strconv.Itoa(result))
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
