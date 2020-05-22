package main

import (
	"fmt"
	"strings"
)

func acroynm (input string) {
	words := strings.Fields(input)
	fmt.Println(words, len(words))
}

func main() {
	acroynm("Ade is eighteen")
	fmt.Println("Done")
}