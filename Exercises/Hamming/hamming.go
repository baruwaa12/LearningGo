package hamming

import (
	"errors"
)

// Distance determines the number of differences between two strings
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, errors.New("lengths are not equal")
	}

	d := 0
	for i := 0; i < len(ar); i++ {
		if ar[i] != br[i] {
			d++
		}
	}

	return d, nil
}