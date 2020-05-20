package sieve

// Sieve find all the primes from 2 up to a given number using the Sieve of Eratosthenes method.
func Sieve(limit int) (primes []int) {
	notInPrimes := make([]bool, limit+1)
	notInPrimes[0] = true
	notInPrimes[1] = true

	for i := 2; i < limit/2; i++ {
		for j := i; i*j <= limit; j++ {
			notInPrimes[i*j] = true
		}
	}

	for i, el := range notInPrimes {
		if !el {
			primes = append(primes, i)
		}
	}
	return primes
}