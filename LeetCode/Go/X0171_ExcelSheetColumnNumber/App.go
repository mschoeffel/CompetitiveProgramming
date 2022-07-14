package X0171_ExcelSheetColumnNumber

import (
	"fmt"
	"strconv"
)

func Main() {
	column := "!"

	result := titleToNumber(column)
	fmt.Println("Result Excel Sheet Column Number: " + strconv.Itoa(result))
}

func titleToNumber(columnTitle string) int {
	number := 0
	for _, i2 := range columnTitle {
		number *= 26
		number += int(i2 - 'A' + 1)
	}
	return number
}
