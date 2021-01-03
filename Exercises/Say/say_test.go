package Say

import (
	"testing"
)

func TestGetHundreds(t *testing.T) {
	for _, test := range getHundredsTestCases {
		if actual := getHundreds(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}

