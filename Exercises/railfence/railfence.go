package main

import (
	"fmt"
)
 



func encoder(rails int, sentence string) (ciphertext string) {
	length := len(sentence)
	for i := 0  ; i <= length; i = i + rails {
		fmt.Println(i)
	}
	return sentence
}

func main() {
	encoder(3, "Hello my name is ade")
}