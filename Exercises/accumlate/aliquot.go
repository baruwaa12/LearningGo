package perfect

import (
	"errors"
	"fmt"
)

var ErrOnlyPositive = ClassificationError{}

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

type ClassificationError struct {
}

func (e ClassificationError) Error() string {
	return fmt.Sprintf("number to classify must be a positive integer")
}

func sumOfFactors(n int64) int64 {
	// Let's try a brute force approach
	var sum int64
	for f := int64(1); f < n; f++ {
		if n%f == 0 {
			sum += f
		}
	}
	return sum
}

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return -1, ErrOnlyPositive
	}
	
	switch v := sumOfFactors(n); {
	case v < n:
		return ClassificationDeficient, nil
	case v == n:
		return ClassificationPerfect, nil
	case v > n:
		return ClassificationAbundant, nil
	default:
		return -1, errors.New("impossible")
	}
}