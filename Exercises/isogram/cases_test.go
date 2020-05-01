package isogram

var testCases = []struct {
	description string
	input   string
	expected bool

}{
	{
		description: "empty string",
		input: "",
		expected: true,
	},
	{
		description: "isogram with only lower case characters",
		input:       "isogram",
		expected:    true,
	},
}