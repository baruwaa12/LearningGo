package reverse

// Reverse function
func Reverse(s string) string {
	r := []rune(s)
	i, n := 0, len(r) - 1
	for i < n {
		r[i], r[n] = r[n], r[i]
		i++
		n--

	}
	return string(r)
}
