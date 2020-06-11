package main

// Factors function finds all prime factors of an integer
func Factors(n int64) []int64 {
	if n == 1 {
		return []int64{}
	}
	factors := make([]int64, 0)

	i := int64(2)
	for ; n%i != 0; i++ {
	}

	factors = append(factors, i)
	subFactors := Factors(n / i)
	factors = append(factors, subFactors...)
	return factors
}

func main() {
	Factors(10)
}