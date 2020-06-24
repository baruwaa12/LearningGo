package main

import (
	"fmt"
	"math/rand"
)


// Creates 4x12 grid
func windowpos(racks int)  {
	rows := 12
	windowrack := make([][]string, rows)

	// loop through racks
	for i := 0; i < rows; i++ {
		 // seed each rack
		rack :=  make([]string, 4)
		rack = seedRackRandomly(rack)

		windowrack[i] = rack
	}
	fmt.Println(windowrack)
}

// Seeds a given rack and populates with one of the 4 keys
func seedRackRandomly(rack []string) []string {
	var keys = []string{"R", "C", "G", "V"}
	for _, item := range rack {
		randomIndex := rand.Intn(3)
		item = keys[randomIndex] 
		fmt.Println(item)
	}
	return rack

}

func main() {
	
}
