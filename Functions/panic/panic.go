package main

import (
	"fmt"
)
func main() {
	f(3)
	fmt.Println("Running...")
}

func f(x int) {
	fmt.Printf"(f(%d)\n", x+0/x) // function panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}