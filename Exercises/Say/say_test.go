package Say

import (
	"testing"
)

func TestNumberChecker(t *testing.T) {
	for _, test := range testCases {
		if actual := numberchecker(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}