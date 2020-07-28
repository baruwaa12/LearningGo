package sieve

import (
	"testing"
	"reflect"
)

// TestSieve tests Sieve Function
func TestSieve(t *testing.T) {
	for _, test := range testCases {
		if actual := sieve(test.limit); reflect.DeepEqual(actual, test.list) == false {
			t.Fatalf("FAILED: %s, %T", test.description, test.limit)
			
		}
		t.Logf("PASSED: %s", test.description)

	}

}