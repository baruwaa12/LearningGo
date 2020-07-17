package main

import (
	"fmt"
)


func buildMatrix(row int, column int) [][]int {

	matrix1 := make([][]int, row)
	for i := 0; i <= column; i++ {
		matrix := make([]int, column)
		fmt.Println(matrix)
	}
	return matrix1
}


func main() {
	buildMatrix(5, 6)
}