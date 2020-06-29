package palindromic

import (
	"testing"
	"reflect"
)


func TestUniqueNumbers(t *testing.T) {
	for _,test := range testCases {
		if actual := unique(test.input); reflect.DeepEqual(actual, test.expected) == false {
			t.Fatalf("FAIL: %s\nInput %q expected [%p], actual [%p]", test.description, test.input, test.expected, actual)
		}
		t.Logf("PASS: %s", test.description)
	}
}

