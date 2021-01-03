package Say

var uptotwenty = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"", "", "twenty", "thirty", "fourty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var hundreds = []string{"", "one-hundred", "two-hundred", "three-hundred", "four-hundred", "five-hundred", "six-hundred", "seven-hundred", "eight-hundred", "nine-hundred"}
var thousands = []string{"", "one-thousand", "two-thousand", "three-thousand", "four-thousand", "five-thousand", "six-thousand", "seven-thousand", "eight-thousand", "nine-thousand "}



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
