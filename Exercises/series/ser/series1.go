package main

import (
	"fmt"
)

func allpossible(n int, s string) []string {

	var firstnumber int
	var substrings []string
	length := len(s)
	fmt.Println(length)

	for end := n; end <= len(s); end++ {
		substrings = append(substrings, s[firstnumber:end])
		fmt.Println(substrings)

	}
	return substrings
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}

func main() {
	allpossible(2, "123764")
	fmt.Println(allpossible)
}
