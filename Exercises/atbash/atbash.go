package main 

import (
	"strings"
	"unicode"
	"fmt"
)

var runeTable = []rune("zyxwvutsrqponmlkjihgfedcba")

// Atbash returns encrypted string using atbash encryption
func Atbash(input string) string {
	input = strings.ToLower(input)
	var result []rune
	counter := 0
	for _, letter := range input {
		if letter >= 'a' && letter <= 'z' {
			result = append(result, runeTable[letter-'a'])
			counter++
		}
		if unicode.IsNumber(letter) {
			result = append(result, letter)
			counter++
		}

		if counter == 5 {
			result = append(result, ' ')
			counter = 0
		}
	}
	fmt.Println(result)
	return strings.TrimSpace(string(result))
}

func main () {
	Atbash("efhizea")
	
}