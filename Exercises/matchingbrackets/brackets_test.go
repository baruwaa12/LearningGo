package matchingbrackets

import (
	"testing"
)

func TestBracket(t *testing.T) {
	for _, tt := range testCases {
		open, close := bracketcounter(tt.input)
		if open != tt.openbrackets {
			t.Fatalf("Bracket(%q) was expected to return %v but returned %v.",
			tt.input, tt.openbrackets, tt.closebrackets)
		} else {
			if close != tt.closebrackets {
				t.Fatalf("Bracket(%q) was expected to return %v but returned %v.",
				tt.input, tt.openbrackets, tt.closebrackets)
			}
		}
	}
}