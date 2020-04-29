package acronym

import (
	"testing"
)

func TestAcronym(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Abbreviate(test.input)
		if actual != test.expected {
			t.Errorf("Acronym test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkAcronym(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Abbreviate(test.input)
		}
	}
}

// goos: windows
// goarch: amd64
// BenchmarkAcronym-4   	   66438	     17799 ns/op	    3354 B/op	      75 allocs/op
// PASS
// ok  	_/c_/OnGithubAde/GoLang/Exercises/acroynm	1.566s
// Success: Benchmarks passed.

// 66438 is the total number of interations completed
// 17799 is the time taken for 1 operation to complete.
// 3