package main

import (
	"fmt"
)

func pascaltriangle(rows int) (triangle [][]int) {
	triangle = make([][]int, rows)
	
	fmt.Println(triangle)
	return triangle
}

func main() {
	pascaltriangle(5)
}