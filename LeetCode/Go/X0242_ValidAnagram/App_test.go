package X0242_ValidAnagram

import (
	"strconv"
	"testing"
)

type InputCase struct {
	s        string
	t        string
	expected bool
	comment  string
}

func TestApp(t *testing.T) {
	testSets := []InputCase{
		{s: "anagram", t: "nagaram", expected: true, comment: "Example1"},
		{s: "rat", t: "car", expected: false, comment: "Example2"},
		{s: "abcdef", t: "abcde", expected: false, comment: "FistLonger"},
		{s: "abcde", t: "abcdef", expected: false, comment: "SecondLonger"},
		{s: "aaa", t: "aab", expected: false, comment: "AmountOne"},
		{s: "a", t: "a", expected: true, comment: "SingleCharHappyCase"},
		{s: "a", t: "b", expected: false, comment: "SingleCharBadCase"},
	}

	for index, element := range testSets {
		t.Run("Test: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testApp(element.s, element.t, element.expected, index, t)
			})

		t.Run("Test Compact: "+strconv.Itoa(index)+"_"+element.comment,
			func(t *testing.T) {
				testAppCompact(element.s, element.t, element.expected, index, t)
			})
	}
}

func testApp(s string, o string, expected bool, index int, t *testing.T) {
	result := isAnagram(s, o)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}

func testAppCompact(s string, o string, expected bool, index int, t *testing.T) {
	result := isAnagramCompact(s, o)
	if result != expected {
		t.Errorf("Error at Testset: " + strconv.Itoa(index) +
			" Expected: " + strconv.FormatBool(expected) +
			" Actual: " + strconv.FormatBool(result))
	}
}

func BenchmarkApp(b *testing.B) {
	benchmarkSets := []InputCase{
		{s: "anagram", t: "nagaram", expected: true, comment: "Example1"},
		{s: "rat", t: "car", expected: false, comment: "Example2"},
		{s: "abcdef", t: "abcde", expected: false, comment: "FistLonger"},
		{s: "abcde", t: "abcdef", expected: false, comment: "SecondLonger"},
		{s: "aaa", t: "aab", expected: false, comment: "AmountOne"},
		{s: "a", t: "a", expected: true, comment: "SingleCharHappyCase"},
		{s: "a", t: "b", expected: false, comment: "SingleCharBadCase"},
		{s: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", t: "aaabbbcccdddeeefffggghhhiiijjjkkklllmmmnnnooopppqqqrrrssstttuuuvvvwwwxxxyyyzzz", expected: true, comment: "LongHappyCase"},
		{s: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", t: "aaabbbcccdddeeefffggghhhiiijjjkkklllmmmnnnooopppqqqrrrssstttuuuvvvwwwxxxyyyzza", expected: false, comment: "LongBadCase"},
	}

	for index, element := range benchmarkSets {
		b.Run("Benchmark: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkApp(element.s, element.t, b) })
		b.Run("Benchmark Compact: "+strconv.Itoa(index)+"_"+element.comment,
			func(b *testing.B) { benchmarkAppCompact(element.s, element.t, b) })
	}
}

func benchmarkApp(s string, o string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagram(s, o)
	}
}

func benchmarkAppCompact(s string, o string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		isAnagramCompact(s, o)
	}
}
