package acronym

import (
	"regexp"
	"strings"
)

// wordRegex matches a words in a string.
var wordRegex = regexp.MustCompile(`([a-zA-Z\']+)`)

// Abbreviate converts a series of words into an acronym.
func Abbreviate(s string) string {
	words := wordRegex.FindAllString(s, -1)

	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(strings.ToUpper(string(w[0])))
	}
	return sb.String()
}