package aliquot

import "testing"



func TestIsPerfect3(t * testing.T) {
	for _, test := range testCases {
		if actual := isPerfect(test.a, test.b, test.c); actual != test.result {
			t.Fatalf("FAIL [%s]", test.description)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func TestIsPerfect4(t * testing.T) {
	for _, test := range testCases {
		if actual := isPerfect4(test.a, test.b, test.c,)
	}
}