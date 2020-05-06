package robot

import (
	"errors"
	"math/rand"
)

const max = 26 * 26 * 1E3


var seenM = make(map[int32]struct{}, max)

// ErrNoMoreName variable
var ErrNoMoreName = errors.New("No More Names")

// Robot type defined
type Robot [5]byte

// Name function
func (r *Robot) Name() (string, error) {
	if r[0] != 0 {
		return string(r[0:5:5]), nil
	}

	if len(seenM) == max {
		return "", ErrNoMoreName
	}

	n := rand.Int31n(max)
	for _, ok := seenM[n]; ok; _, ok = seenM[n] {
		n = rand.Int31n(max)
	}
	seenM[n] = struct{}{}

	for i := 0; i < 2; i++ {
		r[i] = byte(n%26) + 'A'
		n /= 26
	}
	for i := 2; i < 5; i++ {
		r[i] = byte(n%10) + '0'
		n /= 10
	}

	return string(r[0:5:5]), nil
}

// Reset Function
func (r *Robot) Reset() {
	r[0] = 0
}
