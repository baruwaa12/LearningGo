package main

import (
	"fmt"
)

// Grainplacer takes the dimensions of a grid.
// Then squares each board number value starting from one to the end of the board.
func grainplacer(xaxis int, yaxis int) (grain int) {

	// Create a chessboard with x axis
	chessboard := make([][]int, xaxis)
	rowCounter := 0
	for i := 0; i < xaxis; i++ {
		chessboard[i] = make([]int, yaxis)		
		
		// Add the the number of rows with y axis
		for j := 0; j < yaxis; j++ {
			rowCounter++
			fmt.Println(rowCounter)
			cellValue := rowCounter*rowCounter
			chessboard[i][j] = cellValue
		}

	}
	fmt.Println(chessboard)
	return grain
}

func main() {
	grainplacer(4, 2)
}
