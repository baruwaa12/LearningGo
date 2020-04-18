package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Running..")
}

func hypotenuse(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
	
}

func printhypo () {
	fmt.Println(hypotenuse(3, 4)) // "5"
	
}



