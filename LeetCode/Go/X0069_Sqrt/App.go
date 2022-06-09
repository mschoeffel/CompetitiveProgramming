package X0069_Sqrt

import (
	"fmt"
	"sort"
	"strconv"
)

func Main() {
	x := 4

	result := mySqrt(x)
	fmt.Println("Result Sqrt: " + strconv.Itoa(result))
	result = mySqrtSortSearch(x)
	fmt.Println("Result Sqrt Sort Search: " + strconv.Itoa(result))

}

func mySqrt(x int) int {
	l := 0
	r := x
	for l < r {
		mid := (r + l + 1) / 2
		if mid <= x/mid {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return r
}

func mySqrtSortSearch(x int) int {
	v := sort.Search(x, func(i int) bool {
		return i*i >= x
	})
	if v*v > x {
		v--
	}
	return v
}
