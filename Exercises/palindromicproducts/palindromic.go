package palindromic

import (
	"fmt"
)

// Takes a start and end value.
// Loops through 2 seperate arrays and finds all products within the given range
func palinchecker(start int, end int) []int {
	factors := []int{}
	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			factors = append(factors, (i * j))

		}
	}
	return factors
}

// Takes the factors array and removes any duplicate values
// Then creates a seperate array and appends new values
func unique(factors []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range factors {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}


func main() {
	fmt.Println(unique(palinchecker(1, 5)))
}
