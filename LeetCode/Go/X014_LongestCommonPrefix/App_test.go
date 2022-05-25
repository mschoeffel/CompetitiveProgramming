package X014_LongestCommonPrefix

import (
	"reflect"
	"strconv"
	"testing"
)

type InputCase struct {
	words    []string
	expected string
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{words: []string{"flower", "flow", "flight"}, expected: "fl", comment: "Example1"},
		{words: []string{"dog", "racecar", "car"}, expected: "", comment: "Example2"},
		{words: []string{}, expected: "", comment: "EmptyArray"},
		{words: []string{"dog"}, expected: "dog", comment: "ArrayOnlyOneElement"},
		{words: []string{"hello", "hell", "shell"}, expected: "", comment: "OneOff"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testApp(element.words, element.expected, index, t) })
		t.Run("Test FindByFirst: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) { testAppFindByFirst(element.words, element.expected, index, t) })
	}
}

func testApp(roman []string, expected string, index int, t *testing.T) {
	actual := longestCommonPrefix(roman)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}

func testAppFindByFirst(roman []string, expected string, index int, t *testing.T) {
	actual := longestCommonPrefixSearchByFirst(roman)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + expected +
			" Actual: " + actual)
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{words: []string{"flower", "flow", "flight"}, expected: "fl", comment: "Example1"},
		{words: []string{"dog", "racecar", "car"}, expected: "", comment: "Example2"},
		{words: []string{}, expected: "", comment: "EmptyArray"},
		{words: []string{"dog"}, expected: "dog", comment: "ArrayOnlyOneElement"},
		{words: []string{"hello", "hell", "shell"}, expected: "", comment: "OneOff"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.words, b) })
		b.Run("Benchmark FindByFirst: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppFindByFirst(element.words, b) })
	}
}

func benchmarkApp(words []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(words)
	}
}

func benchmarkAppFindByFirst(words []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefixSearchByFirst(words)
	}
}
