package accumlate

func Accumulate(strs []string, f func(x string) string) []string {
	result := make([]string, len(strs))
	for i, v := range strs {
		result[i] = f(v)
	}
	return result
}