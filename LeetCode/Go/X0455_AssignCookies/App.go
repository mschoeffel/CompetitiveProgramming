package X0455_AssignCookies

import (
	"fmt"
	"sort"
	"strconv"
)

func Main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}

	result := findContentChildren(g, s)
	fmt.Println("Result Assign Cookies: " + strconv.Itoa(result))

}

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)
	var out int

	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			out++
			i++
			j++
		} else {
			j++
		}
	}

	return out
}
