package main

import (
	"fmt"
	"strings"
)

// ToRNA translates dna strands into RNA

func toRNA(dna string) (result string) {
	RNA := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	var rna strings.Builder
	for _, c := range dna {
		rna.WriteRune(RNA[c])
	}
	result = rna.string()
	return result
}

func main() {
	toRNA("aacgtttgtaaccag")
	fmt.Println(toRNA)

}
