// Package pythagorean creates lists of pythagorean triplets.
package main

import (
	"fmt"
)

// Triplet is a pythagorean triplet.
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var triplets []Triplet
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				candidate := Triplet{a, b, c}
				// fmt.Println(candidate)
				if isPythagorean(candidate) {
					triplets = append(triplets, candidate)
				}
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(sum int) []Triplet {
	var triplets []Triplet
	for a := 1; a < sum/3; a++ {
		for b := a; b < sum/2; b++ {
			for c := b; c < sum; c++ {
				candidate := Triplet{a, b, c}
				
				if isPythagorean(candidate) && a+b+c == sum {
					triplets = append(triplets, candidate)
				}
			}
		}
	}
	return triplets
}

func isPythagorean(triplet Triplet) bool {
	return triplet[0]*triplet[0]+triplet[1]*triplet[1] == triplet[2]*triplet[2]
}

func main () {
	result:=Range(3, 5)
	fmt.Println(result)
}