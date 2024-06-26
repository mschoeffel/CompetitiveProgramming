package X0482_LicenseKeyFormatting

import (
	"fmt"
	"unicode"
)

func Main() {
	s := "5F3Z-2e-9-w"
	k := 4

	result := licenseKeyFormatting(s, k)
	fmt.Println("Result License Key Formatting: " + result)
}

func licenseKeyFormatting(s string, k int) string {
	runes := []rune(s)
	result := ""
	count := 0

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '-' {
			continue
		}
		result = string(unicode.ToUpper(runes[i])) + result
		count++
		if count%k == 0 && i > 0 {
			result = "-" + result
			count = 0
		}
	}

	if len(result) > 0 {
		resultRunes := []rune(result)
		for resultRunes[0] == '-' {
			resultRunes = resultRunes[1:]
		}
		result = string(resultRunes)
	}
	return result
}
