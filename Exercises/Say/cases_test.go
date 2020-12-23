package Say

var testCases = []struct {
	description string
	input 		int
	expected 	string
}{
	{
		description: "zero",
		input:		 0,
		expected:    "zero", 
	},
	{
		description: "one",
		input:		 1,
		expected:    "one", 
	},	 
	{
		description: "four ",
		input:		 4,
		expected:    "four", 
	},
	{
		description: "twenty ",
		input:		 20,
		expected:    "twenty ", 
	},	
	{
		description: "nineteen",
		input:		 19,
		expected:    "nineteen", 
	},
	{
		description: "fifteen",
		input:		 15,
		expected:    "fifteen", 
	},
	{
		description: "eighty one",
		input:		 81,
		expected:    "eighty one", 
	},
	{
		description: "twenty two",
		input:		 22,
		expected:    "twenty two", 
	}, 
 

}