package X0344_ReverseString

import (
	"encoding/json"
	"fmt"
)

func Main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}

	reverseString(s)
	output, _ := json.Marshal(s)
	fmt.Println("Result Reverse String: " + string(output))

}

func reverseString(s []byte) {
	if len(s) == 1 {
		return
	}

	for i := 0; i < len(s)/2; i++ {
		save := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = save
	}
}
