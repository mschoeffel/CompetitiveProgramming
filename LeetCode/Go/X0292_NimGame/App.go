package X0292_NimGame

import (
	"fmt"
	"strconv"
)

func Main() {
	num := 4

	output := canWinNim(num)
	fmt.Println("Result Nim Game: " + strconv.FormatBool(output))

}

func canWinNim(n int) bool {
	return !(n%4 == 0)
}
