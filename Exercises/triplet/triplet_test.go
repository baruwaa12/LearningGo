package triplet

import (
	"testing"
)

func TestTripletChecker(t *testing.T) {
	for _, test := range PythagoreanTriplets {
		var result = tripletchecker(test.a, test.b, test.c);
		if result != test.total {
			t.Fatalf("FAILED given a %v b %v c %v returned %v expected %v", test.a, test.b, test.c, result, test.total)
		}
		t.Logf("Passed %s", test.description)
	}
}

