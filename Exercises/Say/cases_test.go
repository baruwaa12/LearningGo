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
