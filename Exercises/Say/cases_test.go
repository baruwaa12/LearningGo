package Say

var getHundredsTestCases = []struct {
	description string
	input 		int
	expected 	string
}{
	{
		description: "one hundred",
		input:		 100,
		expected:    "one-hundred", 
	},
	{
		description: "nine hundred and ninety nine",
		input:		 999,
		expected:    "nine-hundred", 
	},
	{
		description: "ninety nine",
		input:		 99,
		expected:    "", 
	},
	{
		description: "four hundred and fifty",
		input:		 450,
		expected:    "four-hundred", 
	},
}


var getTensTestCase = []struct {
	description string
	input 		int
	expected 	string
}{
	{
		description: "ten",
		input:		 10,
		expected:    "", 
	},
	{
		description: "twenty",
		input:		 20,
		expected:    "twenty", 
	},
	{
		description: "no tens",
		input:		 100,
		expected:    "", 
	},
	{
		description: "one thousand one hundred and thirty",
		input:		 130,
		expected:    "thirty", 
	},
	{
		description: "four hundred and fifty",
		input:		 450,
		expected:    "fifty", 
	},
}


var getUnitsTestCases = []struct {
	description string
	input 		int
	expected 	string
}{
	{
		description: "ten",
		input:		 10,
		expected:    "", 
	},
	{
		description: "two hundred and fifty one",
		input:		 251,
		expected:    "one", 
	},
	{
		description: "sixty one",
		input:		 61,
		expected:    "one", 
	},
	{
		description: "one thousand one hundred and thirty five",
		input:		 135,
		expected:    "five", 
	},
	{
		description: "nine hundred and ninety nine",
		input:		 999,
		expected:    "nine", 
	},
}


var getFinalNumberTestCases = []struct {
	description string
	input 		int
	expected 	string
}{
	{
		description: "one hundred and one",
		input:		 101,
		expected:    "one-hundred and one", 
	},
	{
		description: "one hundred and twenty one",
		input:		 121,
		expected:    "one-hundred and twenty one", 
	},
	{
		description: "five hundred and twenty ",
		input:		 520,
		expected:    "five-hundred and twenty ", 
	},
	{
		description: "five hundred and twenty ",
		input:		 520,
		expected:    "five-hundred and twenty ", 
	},
	{
		description: "nine hundred and sixty ",
		input:		 960,
		expected:    "nine-hundred and sixty  ", 
	},
}
