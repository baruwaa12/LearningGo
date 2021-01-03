package Say

var uptotwenty = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "", "twenty ", "thirty ", "fourty ", "fifty ", "sixty ", "seventy ", "eighty ", "ninety"}
var hundreds = []string{"", "one-hundred", "two-hundred", "three-hundred", "four-hundred", "five-hundred", "six-hundred", "seven-hundred", "eight-hundred", "nine-hundred"}
var thousands = []string{"", "one-thousand ", "two-thousand", "three-thousand ", "four-thousand ", "five-thousand ", "six-thousand ", "seven-thousand ", "eight-thousand ", "nine-thousand "}



func getHundreds(number int) (string) {
	NumOfHundreds := number / 100
	if NumOfHundreds >= 1  || NumOfHundreds < 9 {
		return hundreds[NumOfHundreds]
	}
	return ""
}

func getTens(number int) (numberoftens string) {
	NumOfTens := number / 10
	if NumOfTens >= 1  || NumOfTens < 9 {
		return tens[NumOfTens]
	}
	return ""
}

func getUnits(number int) (numberofunits string) {
	return ""
}
