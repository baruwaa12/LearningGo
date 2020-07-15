package piglatin

import (
	"fmt"
	"strings"
)


func engtoPigLatin(word string) string {
	result := ""
	wordLower := strings.ToLower(word)
	if wordLower[0] == 'a' || wordLower[0] == 'e' || wordLower[0] == 'i' || wordLower[0] == 'o' || wordLower[0] == 'u' {
		word = word + "ay"
		result = word
		fmt.Println(result)
	if wordLower[0][1] == "ay" {
		
	}

	}
	return result
}