package X0482_LicenseKeyFormatting

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	k        int
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "5F3Z-2e-9-w", k: 4, expected: "5F3Z-2E9W", comment: "Example1"},
		{s: "2-5g-3-J", k: 2, expected: "2-5G-3J", comment: "Example2"},
		{s: "E3gh527i", k: 2, expected: "E3-GH-52-7I", comment: "NoDashes"},
		{s: "E-3-g-h-5-2-7-i", k: 8, expected: "E3GH527I", comment: "OneSegment"},
		{s: "E-3-g-h-5-2-7-i", k: 1, expected: "E-3-G-H-5-2-7-I", comment: "MinSegment"},
		{s: "Er4-gp0-56e", k: 3, expected: "ER4-GP0-56E", comment: "AlreadyFormatted"},
		{s: "e", k: 1, expected: "E", comment: "SingleRune"},
		{s: "e", k: 5, expected: "E", comment: "KBiggerThanSingleRune"},
		{s: "--a-a-a-a--", k: 2, expected: "AA-AA", comment: "LeadingDashes"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.k, element.expected, index, t)
			})
	}
}

func testApp(s string, k int, expected string, index int, t *testing.T) {
	actual := licenseKeyFormatting(s, k)
	if actual != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}
