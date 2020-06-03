package main

import (
	"strings"
	"fmt"
)

// ToRNA translates dna strands into RNA

func toRNA(dna string) string {
	 RNA := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	 var rna strings.Builder
	for _, c := range dna {
		rna.WriteRune(RNA[c])
	}

	return rna.String()
}

func main() {
	toRNA("aacgtttgtaaccag")

	
}
