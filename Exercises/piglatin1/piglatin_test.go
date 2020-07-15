package piglatin

import (
	"testing"
	"reflect"
)


func TestLetters(t *testing.T) {
	for _,test := range testCases {
		if actual := engtoPigLatin(test.word); reflect.DeepEqual(actual, test.translated) == false {
			t.Fatalf("FAIL: %s\nInput %q types [%s], actual [%s]", test.description, test.word, test.translated, actual)
		}
		t.Logf("PASS: %s", test.description)
	}
}