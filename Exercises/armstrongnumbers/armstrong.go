package main

import (
	"fmt"
	// "math"
	"strconv"
)

// armstrongchecker returns if any given number is an armstrong number
func armstrongchecker(number string) (result bool) {
	str1 := number
	i1, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println(i1)
	}
	if i1^(len(str1)) != i1 {
		result = false
	}
	return result
}


func main() {
	armstrongchecker("9")
}