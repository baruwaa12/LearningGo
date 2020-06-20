package main

import (
	"fmt"
)


// Create a function to check if a grid contains the number 1
func isNumberOne(grid [4][3]string) bool {

	var i, j int

	// Rule 1 checks if position (1,1) and (1,2) have pipes if no it isnt a 1
	if(grid[1][1] != "|" ||  grid[1][2]  != "|"){
		// Not 1
		return false
	}

	// Rule 2 checks if all positions in the grid apart from (1,1) and (1,2) are empty
	for i = 0; i <= 3; i++ {
		fmt.Println("")
		for j = 0; j <= 2; j++ {
			if( i == 1 && (j == 1 || j == 2) ){
				// ignore
			
			}else{
				if(grid[i][j] != "") {
					// Not 1
					return false
				}
			}
		}
	}
	return true
}


func main() {

	// Create a 3x4 grid with random symbols to test
	var gridNotNumberOne = [4][3]string{{"|", "_", "?"}, {"|", "_", "?"}, {"|", "_", "?"}, {"|", "_", "?"}}

	// Create an empty grid with 3x4 dimensions
	var gridWithnumberOne = [4][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}, {"", "", ""}}

	// Insert pipes at position (1,1) and (1,2)
	gridWithnumberOne[1][1] = "|"
	gridWithnumberOne[1][2] = "|"

	// test should return false
	fmt.Println(isNumberOne(gridNotNumberOne)) 

	// test should return true
	fmt.Println(isNumberOne(gridWithnumberOne)) 

}
