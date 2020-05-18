package pangram

import (
	"strings"
	"unicode"
)

// IsPangram function
func IsPangram(input string) bool {
	bs := [26]bool{}
	for _, r := range strings.ToLower(input) {
		if !unicode.IsLetter(r) { continue }
		bs[int(r - 'a')] = true
	}
	for _, b := range bs {
		if !b { return false }
	}
	return true
}