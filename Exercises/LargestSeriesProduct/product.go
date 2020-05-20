package main

import (
	// "bufio"
	"fmt"
	"strconv"
	// "os"
)

func collectnumber() {
	// number := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your number: ")
	// number, _ := reader.ReadString('\n')
}

// Calculate the length of the number
// Set the values of current and largest so they can be reset outside of each iteration
// of any span length
// Loop through the number string given the span length
// Calculate the largest and current value then compare and store the largest 
// Continue until loop is completed.

func ndecider(mynumber string, span int) (int) {
	fmt.Println(len(mynumber))

	current := 1
	largest := 1


	for i := 0; i < len(mynumber)-span; i++ {
		current = 1
		for z := 0; z < span; z++ {
			a, _ := strconv.Atoi(string(mynumber[i+z]))
			if i == 0 {
				largest = largest * a
			} 
			current = current * a
		}

		if largest < current{
			largest = current
	
		}

		fmt.Println("-----------")

	}
	return largest
}

func main() {
	fmt.Println("Done")
	result := ndecider("102330", 2)
	fmt.Println(result)
}
