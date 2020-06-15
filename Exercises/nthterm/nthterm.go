 //Package prime provdes the nth prime factor
package prime

//Nth returns the nth prime number
func Nth(n int) (int, bool) {
	prime := []int{2}
	if n == 0 {
		return 0, false
	}
	if n == 1 {
		return 2, true
	}
	count := 1
	i := 3
	for count < n {
		status := true
		for j := 2; j < (i/2)+1; j++ {
			if i%j == 0 {
				status = false
				break
			}
		}
		if status == true {
			prime = append(prime, i)
			count++
		}

		i++
	}
	return 
}