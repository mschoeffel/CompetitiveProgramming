package X0278_FirstBadVersion

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 5

	output := firstBadVersion(num)
	fmt.Println("Result First Bad Version: " + strconv.Itoa(output))

}

func firstBadVersion(n int) int {
	lower := 0
	upper := n
	diff := upper - lower
	for diff > 1 {
		check := lower + (diff / 2)
		version := isBadVersion(check)
		if version {
			upper = check
		} else {
			lower = check
		}
		diff = upper - lower
	}
	return upper
}

func isBadVersion(version int) bool {
	return version == 4
}
