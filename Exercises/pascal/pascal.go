package pascal

func Triangle(n int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			arr[i][j] = Fact(i) / (Fact(i-j) * Fact(j))
		}
	}
	return arr
}

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}