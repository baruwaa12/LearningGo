package sumofmultiples

import "testing"

func TestSumMultiples(t *testing.T) {
	for _, test := range varTests {
		s := multipleschecker(test.limit, test.number)
		if s != test.sum {
			t.Fatalf("Sum of multiples of %v to %d returned %d, want %d.",
				test.number, test.limit, s, test.sum)
		}
	}
}
