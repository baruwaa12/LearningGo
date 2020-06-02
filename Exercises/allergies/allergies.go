package main

import (
	"fmt"
)

var food = []string {"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}



func allergies(total int) (results []string) {
	output := make([]string, 0)
	for i, f := range food {
		if results&(1<<i) == (1 << i) {
			output = append(output, f)
		}
	}
	return output
}


// AllergicTo checks for the allergic condition.
func AllergicTo(n uint, s string) bool {
	for i, f := range food {
		if f == s {
			return n&(1<<i) > 0
		}
	}
	return false
}