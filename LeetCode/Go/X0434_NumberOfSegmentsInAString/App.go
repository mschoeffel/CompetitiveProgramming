package X0434_NumberOfSegmentsInAString

import (
	"fmt"
	"strconv"
)

func Main() {
	s := "Hello, my name is John"

	result := countSegments(s)
	fmt.Println("Result Number of Segments in a String: " + strconv.Itoa(result))

}

func countSegments(s string) int {
	count := 0
	w := false
	for _, v := range s {
		if v != ' ' {
			w = true
		} else {
			if w {
				count++
			}
			w = false
		}
	}
	if w {
		count++
	}
	return count
}
