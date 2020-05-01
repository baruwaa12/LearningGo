package isogram

import (
	"strings"
	"unicode"
	)

// IsIsogram returns true if string is isogam
func IsIsogram(s string) bool {
	lower := strings.ToLower(s)

	set := make(map[rune]bool)
	for _, r := range lower {
		if !unicode.IsLetter(r) {
			continue
		}

		if set[r] {
			return false
		
		}
		set[r] = true
	
	}

	return true
}