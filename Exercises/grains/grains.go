package grains

import (
	"errors"
)

// Square returns the square 
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("the number should be between 1 and 64, inclusive")
	}
	return 1 << (n - 1), nil;
}

// Total returns the total number of grains
// Soo young, so baad! :)
func Total() uint64 {
	return 18446744073709551615
}