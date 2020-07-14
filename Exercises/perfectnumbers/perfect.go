package perfectnumber

import (
	"fmt"
)

func perfectaliquot(number int) (types string) {
	var sum int
	for i := 1; i < number; i++ {
		if number % i == 0 {
			sum += i
			fmt.Println(sum)
		}
		if sum > number {
			types = "abundant"
		} else {
			if sum == number {
				types = "perfect"
			} else {
				if sum < number {
					types = "deficient"
				}
			}

		}

	}
	return types
}
