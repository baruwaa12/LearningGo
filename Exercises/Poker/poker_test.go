package poker

import (
	"testing"
	"reflect"
)

// TestBestHandValid function checks for the best hand in the deck

func TestBestHandValid(t *testing.T) {
	for _, tt := range validTestCases {
		actual, err := BestHand(tt.hands)
		if err != nil {
			var _ error = err
			t.Fatalf("Got unexpected error in valid case %q: %v", tt.name, err)
		}
		if !reflect.DeepEqual(actual, tt.best) {
			t.Fatalf("Mismatch in result of valid case %q: got %#v, want %#v",
				tt.name, actual, tt.best)
		}
		t.Logf("PASSED")
	}
}
