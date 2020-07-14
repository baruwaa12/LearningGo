package perfectnumber

import (
	"testing"
	"reflect"
)


func TestPerfectNumbers(t *testing.T) {
	for _,test := range testCases {
		if actual := perfectaliquot(test.number); reflect.DeepEqual(actual, test.types) == false {
			t.Fatalf("FAIL: %s\nInput %q types [%p], actual [%p]", test.description, test.number, test.types, actual)
		}
		t.Logf("PASS: %s", test.description)
	}
}
