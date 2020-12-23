package Say

import (
	"fmt"
)

var uptotwenty = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "", "twenty", "thirty ", "fourty ", "fifty ", "sixty ", "seventy ", "eighty ", "ninety "}

func numberchecker(number int) (numberinwords string) {
	if number <= 19 {
		numberinwords = uptotwenty[number]
		fmt.Println(numberinwords)
		return numberinwords
	}
	if number <= 90 {
		index := (number/10)
		numberinwords = tens[index]
		fmt.Println(numberinwords)
		return numberinwords
		}
	return numberinwords
}

func main() {
	numberchecker(70)
}
