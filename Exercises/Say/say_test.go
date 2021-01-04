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

func TestGetTens(t *testing.T) {
	for _, test := range getTensTestCase {
		if actual := getTens(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}

func TestGetUnits(t *testing.T) {
	for _, test := range getUnitsTestCases {
		if actual := getUnits(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}

func TestGetFinalNumber(t *testing.T) {
	for _, test := range getFinalNumberTestCases {
		if actual := joinNumbers(test.input, ""); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}

func TestPostFixJoiner(t *testing.T) {
	for _, test := range PostFixJoinerTestCases {
		if actual := NumberPostFixJoiner(test.input); actual != test.expected {
			t.Errorf("Convert(%d) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}
