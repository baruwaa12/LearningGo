package Say

import (
	"fmt"
	"strconv"
)



var uptotwenty = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "", "twenty ", "thirty ", "fourty ", "fifty ", "sixty ", "seventy ", "eighty ", "ninety "}
var hundreds = []string{"", "one-hundred", "two-hundred", "three-hundred", "four-hundred", "five-hundred", "six-hundred", "seven-hundred", "eight-hundred", "nine-hundred"}
var thousands = []string{"", "one-thousand", "two-thousand", "three-thousand", "four-thousand", "five-thousand", "six-thousand", "seven-thousand", "eight-thousand", "nine-thousand"}
var postfix = []string{"", "thousand", "million", "billion", "trillion"}


func getHundreds(number int) (string) {
	NumOfHundreds := number / 100
	if NumOfHundreds >= 1  || NumOfHundreds < 9 {
		return hundreds[NumOfHundreds]
	}
	return ""
}

func getTens(number int) (string) {
	
	NumOfTens := 0

	if number >= 100 {
		NumOfHundreds := number/100
		NumOfTens = (number - (100*NumOfHundreds)) / 10
	}else{
		NumOfTens = number / 10
	}

	if NumOfTens >= 2  || NumOfTens <= 9 {
		return tens[NumOfTens]
	}
	return ""
}

func getUnits(number int) (string) {

	Units := 0
	NumOfHundreds := number/100
	NumOfTens := (number - (100*NumOfHundreds)) 
	Units = NumOfTens - ((NumOfTens/10)*10)

	if Units < 1 || Units > 9 {
		return ""
	}

	return uptotwenty[Units]

}

func joinNumbers(number int, postfix string) (string) {
	finalNumber:= ""
	if number > 100{
		finalNumber = getHundreds(number) + " and " + getTens(number) + "" + getUnits(number)
	}else{
		finalNumber = getTens(number) + "" + getUnits(number)
	}
	return finalNumber + " " + postfix
	
	
}



func NumberPostFixJoiner(number int) (string) {
	finalDigitInWords := ""
	numberString := strconv.Itoa(number)
	postfixIndex := -1
	for i := len(numberString) - 1; i >= 0; i = i - 3 {
		digits := ""
		postfixIndex += 1
		if i < 3 {
			digits = numberString[0:i+1]
		}else{
			digits = numberString[i-2: i+1]
		}																			 
		fmt.Println(digits)
		digitnumb, _ := strconv.Atoi(digits)
		digitInWords := joinNumbers(digitnumb, postfix[postfixIndex])
		finalDigitInWords = digitInWords + " " + " " + finalDigitInWords

	}
	return finalDigitInWords
}



