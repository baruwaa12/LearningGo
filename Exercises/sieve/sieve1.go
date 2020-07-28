package sieve


func sieve(n int) []int {
	if n < 2 {
		return []int{}
	}

	primes := []int{2}

	for i := 3; i <= n; i++ {
		hasPrimeFactor := false
		for j, last := 0, len(primes); j < last; j++ {
			if i%primes[j] == 0 {
				hasPrimeFactor = true
				break
			}
		}
		if !hasPrimeFactor {
			primes = append(primes, i)
		}
	}
	return primes
}


