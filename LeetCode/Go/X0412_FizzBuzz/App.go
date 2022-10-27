package X0412_FizzBuzz

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func Main() {
	num := 3

	result := fizzBuzz(num)
	output, _ := json.Marshal(result)
	fmt.Println("Result Fizz Buzz: " + string(output))

}

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}
