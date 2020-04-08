package main

import (
	   "fmt"
	   "os"
)
func helloworld() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "hello world"
	}
	fmt.Println(s)
}



