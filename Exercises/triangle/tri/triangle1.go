package main

import (
	"fmt"
)

const (
	// Equilaterial triangle
	equal string = "Equilaterial"
	// Iso celes triangle
	iso string = "Isoceles"
	// Sca triangle
	Sca string = "Scalene"
	// NaT a triangle
	NaT string = "Not a triangle"
)


func decider(sideA, sideB, sideC int) (kind string) {
	if sideA == sideB && sideB == sideC && sideA == sideC {
		kind = equal
		fmt.Println(kind)
	} else if (sideA == sideB) && sideB != sideC {
		kind = iso
		fmt.Println(kind)
	} else if (sideA != sideB) && sideB != sideC {
		kind = Sca
		fmt.Println(kind)
	} else if (sideA == 0) && sideB != sideC {
		kind = NaT
		fmt.Println(kind)
	}
	return kind
}

func main() {
	decider(0, 3, 4)
	fmt.Println(decider)
}
