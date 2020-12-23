package aliquot

import "errors"

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("Invalid")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 1, ErrOnlyPositive
	}

	var fact []int64
	var i int64
	for i = 1; i < n; i++ {
		if n%i == 0 {
			fact = append(fact, int64(i))
		}
	}
	var sum int64
	for _, val := range fact {
		sum += val
	}

	var classification Classification
	switch {
	case n == sum:
		classification = 0
	case n < sum:
		classification = 1
	case n > sum:
		classification = 2
	}
	return classification, nil
}