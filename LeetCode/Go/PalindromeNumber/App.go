package PalindromeNumber

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 121

	result := palindromeNumber(num)
	fmt.Println("Result Palindrome Number: " + strconv.FormatBool(result))

	result = palindromeNumberReverseNumber(num)
	fmt.Println("Result Palindrome Number ReverseNumber: " + strconv.FormatBool(result))
}

func palindromeNumber(num int) bool {
	numAsArray := []rune(strconv.Itoa(num))
	size := len(numAsArray)

	for i := 0; i < size/2; i++ {
		numberFromFront := numAsArray[i]
		numberFromBack := numAsArray[size-1-i]
		if numberFromFront != numberFromBack {
			return false
		}
	}
	return true
}

func palindromeNumberReverseNumber(num int) bool {
	var reversed int
	tmp := num
	for tmp > 0 {
		reversed = reversed*10 + tmp%10
		tmp = tmp / 10
	}
	return num == reversed
}
