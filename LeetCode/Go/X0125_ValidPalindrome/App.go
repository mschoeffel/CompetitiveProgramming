package X0125_ValidPalindrome

import (
	"fmt"
	"strconv"
)

func Main() {
	input := "A man, a plan, a canal: Panama"

	result := isPalindrome(input)
	fmt.Println("Result Valid Palindrome: " + strconv.FormatBool(result))

}

func isPalindrome(s string) bool {
	pointerLeft := 0
	pointerRight := len(s) - 1

	for pointerLeft <= pointerRight {
		for pointerLeft < pointerRight && !((s[pointerLeft] >= 'A' && s[pointerLeft] <= 'Z') || (s[pointerLeft] >= 'a' && s[pointerLeft] <= 'z') || (s[pointerLeft] >= '0' && s[pointerLeft] <= '9')) {
			pointerLeft++
		}
		for pointerRight > pointerLeft && !((s[pointerRight] >= 'A' && s[pointerRight] <= 'Z') || (s[pointerRight] >= 'a' && s[pointerRight] <= 'z') || (s[pointerRight] >= '0' && s[pointerRight] <= '9')) {
			pointerRight--
		}
		leftChar := s[pointerLeft]
		rightChar := s[pointerRight]
		if leftChar >= 'A' && leftChar <= 'Z' {
			leftChar += 'a' - 'A'
		}
		if rightChar >= 'A' && rightChar <= 'Z' {
			rightChar += 'a' - 'A'
		}
		if leftChar != rightChar {
			return false
		}
		pointerLeft++
		pointerRight--

	}
	return true
}
