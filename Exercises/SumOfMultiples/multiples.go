package sumofmultiples


// Takes a number and finds all multiples up to a limit and calculates the sum
func multipleschecker(number int, limit int) (sum int) {

	for i := 1; i <= limit; i++ {
		if (number % i) == 0 {
			sum += i
		}
	}
	return sum
}
