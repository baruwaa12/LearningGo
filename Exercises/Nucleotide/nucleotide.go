// Adenine
// Cytosine
// Guanine
// Thymine

package nucleotide

import (
	"errors"
)

type counter map[string]int

// dnachecker takes a DNA strand and counts all frequencies of proteins which
// make up DNA.

func dnachecker(dnaStrand string) (counter, error) {

	// If the string is empty all proteins values should be 0
	frequency := counter{"A": 0, "C": 0, "G": 0, "T": 0}
	if dnaStrand == "" {
		return frequency, nil
	}
	for _, chr := range dnaStrand {
	
		// Using a switch statement look for all proteins and increase the frequency by 1 each time any
		// protein is seen.
		// If the strand contains any other values apart from the 4 proteins
		// return Invalid DNA string.
		switch strand := chr; strand {
		case 'A':
			frequency["A"] = frequency["A"] + 1
		case 'C':
			frequency["C"] = frequency["C"] + 1
		case 'G':
			frequency["G"] = frequency["G"] + 1
		case 'T':
			frequency["T"] = frequency["T"] + 1
		default:
			return frequency, errors.New("Invalid DNA strand")
		}

	}
	return frequency, nil
}
