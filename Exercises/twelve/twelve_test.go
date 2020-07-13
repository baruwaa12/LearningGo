package twelve

import (
	"testing"
)
func TestVerse(t *testing.T) {
	for _, test := range testCases {
		actual := twelvedayschristmaslyrics(test.input)
		if actual != test.expected {
			t.Errorf("Twelve Days test [%d], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}