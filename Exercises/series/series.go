package series

import (
	"strings"
)

// All function to define the substring of length n out of a string s **/
func All(n int, s string) []string {
	length := len(s)
	result := make([]string, 0, length)
	if n > len(s) {
		return nil
	}
	for i := 0; i < length; i++ {
		if i+n <= length {
			val := s[i : i+n]
			if val != "" {
				result = append(result, strings.TrimSpace(val))
			}
		}
	}

	return result
}

// UnsafeFirst function gets the first substring of length n out of s
func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[0:n]
}

// First function gets the first safe substring
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return UnsafeFirst(n, s), true
}