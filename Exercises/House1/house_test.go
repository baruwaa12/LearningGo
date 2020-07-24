package house1

import (
	"testing"
	"reflect"
)

func TestLyricSpitter(t *testing.T) {
	for _, test := range testCases {
		if actual := lyricSpitter(test.noun, test.action, test.oldVerse); reflect.DeepEqual(actual, test.newVerse) == false {
			t.Fatalf("FAILED: %s Input %s expected, actual[%s] newVerse[%T]", test.noun, test.action, test.newVerse, actual)
		}
		t.Logf("PASSED %s %s", test.noun, test.newVerse)
	}
}



func TestReplaceThisIsWithThatAndAction(t *testing.T) {
	for _, test := range replaceThisIsWithThatAndActionTestCases {
		actual := replaceThisIsWithThatAndAction(test.line, test.action);
		if actual != test.newLine {
			t.Fatalf("FAILED: Inputs: '%s ; %s ',  Expected: '%s' Actual: '%s'  ", test.line, test.action, test.newLine, actual)
		}
		t.Logf("PASSED: Inputs: '%s ; %s' Result: '%s' ",  test.line, test.action, actual)
	}
}


func TestNewLine(t *testing.T) {
	for _, test := range newLineTestCases {
		actual := newLine(test.noun);
		if actual != test.newLine {
			t.Fatalf("FAILED: Inputs: '%s',  Expected: '%s' Actual: '%s'  ", test.noun, test.newLine, actual)
		}
		t.Logf("PASSED: Inputs: '%s' Result: '%s' ",  test.noun, actual)
	}
}


