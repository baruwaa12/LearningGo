package nucleotide

import (
	"reflect"
	"testing"
)

func TestCounts(t *testing.T) {
	for _, tc := range testCases {
		dna, err := dnachecker(tc.strand)
		switch {
		case tc.errorExpected:
			if err == nil {
				t.Fatalf("FAIL: %s\nCounts(%q)\nExpected error\nActual: %#v",
					tc.description, tc.strand, dna)
			}
		case err != nil:
			t.Fatalf("FAIL: %s\nCounts(%q)\nExpected: %#v\nGot error: %q",
				tc.description, tc.strand, tc.expected, err)
		case !reflect.DeepEqual(dna, tc.expected):
			t.Fatalf("FAIL: %s\nCounts(%q)\nExpected: %#v\nActual: %#v",
				tc.description, tc.strand, tc.expected, err)
		}
		t.Logf("PASS: %s", tc.description)
	}
}