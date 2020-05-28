package main

import (
	"fmt"
)

func All(space int, str string) []string {
	var substrings []string

	for i := 0; i <= len(str) {
		for z := i+1; i <= len(str) {
			value = i + z
			fmt.Println(value)
		}
}
}

func main() {
	All(2, "123456")
}