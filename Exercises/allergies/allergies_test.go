package allergies

import (
	"testing"
)

func TestAllergies(t *testing.T) {
	for _, test := range Allergies {
		result := test.score
		if result != test.score {
			t.Fatalf("FAILED given score")
		}
		t.Logf("PASSED")

	}
}
