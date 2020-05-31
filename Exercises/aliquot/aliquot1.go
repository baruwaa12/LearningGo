package main

import (
	"fmt"
)

func isPerfect(a int, b int, c int) (result bool) {
	if a+b != c {
		result = false
		fmt.Println(result)
		return result
	} else {
		if a+b == c {
			result = true
			fmt.Println(result)

		}

	}
	return result
}

func main() {
	result := isPerfect(1, 2, 3)
	fmt.Println(result)
}
