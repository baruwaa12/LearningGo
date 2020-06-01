package anagram

import (
	"strings"
)

var letters = map[rune]int{
	'a': 2, 'b': 3, 'c': 5, 'd': 7, 'e': 11, 'f': 13, 'g': 17, 'h': 19,
	'i': 23, 'j': 29, 'k': 31, 'l': 37, 'm': 41, 'n': 43, 'o': 47, 'p': 53,
	'q': 59, 'r': 61, 's': 67, 't': 71, 'u': 73, 'v': 79, 'w': 83, 'x': 89,
	'y': 97, 'z': 101,
}

// isAnagram check whether given strings are anagrams or not
func isAnagram(given, toCheck string) bool {
	given, toCheck = strings.ToLower(given), strings.ToLower(toCheck)
	if given == toCheck {
		return false
	}
	if len(given) != len(toCheck) {
		return false
	}
	prod1, prod2 := 1, 1
	for i, c := range given {
		prod1 *= letters[c]
		prod2 *= letters[rune(toCheck[i])]
	}

	if prod1 != prod2 {
		return false
	}
	return true
}

// Detect returns a list of anagrams for given input from given list
func Detect(input string, list []string) []string {
	ret := make([]string, 0)
	for _, c := range list {
		if isAnagram(input, c) {
			ret = append(ret, c)
		}
	}
	return ret
}