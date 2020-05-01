package reverse

type reverseTestCase struct {
	description string
	input  string
	expected string
}

var testCases = []reverseTestCase{
	{
		description: "an empty string",
		input: "",
		expected: "",
	},
	{
		
		description: " a word",
		input: "robot",
		expected:  "tobor",

	},
	
}
