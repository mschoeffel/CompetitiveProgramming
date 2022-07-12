package X0168_ExcelSheetColumnTitle

import (
	"fmt"
)

func Main() {
	column := 1

	result := convertToTitle(column)
	fmt.Println("Result Excel Sheet Column Title: " + result)
}

func convertToTitle(columnNumber int) string {
	title := ""
	for columnNumber > 26 {
		reminder := columnNumber % 26
		if reminder == 0 {
			title = "Z" + title
			columnNumber = (columnNumber / 26) - 1
		} else {
			title = string(rune(reminder-1+'A')) + title
			columnNumber = columnNumber / 26
		}
	}
	title = string(rune('A'+columnNumber-1)) + title
	return title
}
