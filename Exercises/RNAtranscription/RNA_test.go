package rna 

import (
	"testing"
)

func TestRna(t *testing.T)  {
	for _, test := range testCases {
		if actual := toRNA(test.dna); actual != test.result {
			t.Fatalf("FAILED: %s", test.dna)
		}
		t.Logf("PASSED: %s", test.description)
	}
}