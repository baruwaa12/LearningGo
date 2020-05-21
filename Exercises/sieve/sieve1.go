package main

import (
	"fmt"
)

func sieve(limit int) (list []int) {
	isNotPrimes := make([]int, limit-1)
	for i := 2; i <= limit; i++ {
		if i%2 == 0 && i > 2 {
			isNotPrimes[i-2] = 0
		} else {
			isNotPrimes[i-2] = i
		}
	}

	fmt.Println(isNotPrimes)

	for i := 3; i <= limit; i+=3 {
		fmt.Println(i)
		if i > 3 {
			isNotPrimes[i-2] = 0
		}
	}
	fmt.Println(isNotPrimes)

	primes := make([]int, limit-1)
	// for i := 2; i < len(isNotPrimes); i++ {
	// 	if !isNotPrimes[i] {
	// 		primes = append(primes, i)
	// 	}
	// }
	return primes
}

func main() {
	sieve(9)
	println("Done")
}
