package main

import (
	"fmt"
)

func palindromechecker(number string) bool {
	if len(number) == 3 {
		if int(number[0]) == int(number[2]) {
			return true
		}
		if len(number) == 4 {
			length := int(len(number))
			oppositeside := int(number[length-1])
			if int(number[0]) != int(oppositeside-length-1) {
				return false
			} else {
				return true
				fmt.Println("Number is a palindrome")
			}

		}
	}
	return true
}

func main() {
	palindromechecker("101")
	fmt.Println(palindromechecker)
}
