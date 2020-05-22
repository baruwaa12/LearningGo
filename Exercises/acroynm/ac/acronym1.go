package main

import (
	"fmt"
	"strings"
)


func acroynm (input string) (string) {
	words := strings.Fields(input)
	fmt.Println(words, len(words))
	var result string
	for _, word := range words {
		// fmt.Println(string(word[0]))
		result = result + string(word[0])
		// fmt.Println(result)
	}
	return result
}

func main() {
	fmt.Println(acroynm("Ade is eighteen"))
	fmt.Println("Done")
}