package saddlepoints

import (
	"fmt"
)

// [0 0 0 0]
// [0 0 0 0]
// [0 0 0 0]
// [0 0 0 0]
// Build a 2d matrix defaulted with values 0
// 2D matrix is presented as a 2D array
func buildMatrix(row int, column int) [][]int {
	
	// Create a 2D slice with a number of rows and columns.
	matrix1 := make([][]int, row)

	// Set the default values of each inner array to 0.
	for i := 0; i < column; i++ {
		
		// Create an inner array inside the outer array.
		// Looping through to the number of columns and set each inner array value
		// to 0.
		inner := make([]int, column)
		for j := 0; j < column; j++ {
			inner[j] = 0
		}
		matrix1[i] = inner
		fmt.Println(inner)
	}
	return matrix1
}

