package matchingbrackets

type bracketTest struct {
	input    string
	openbrackets int
	closebrackets int
}

var testCases = []bracketTest {
	{
		// pair of parenthesis
		"(())",
		2,
		2,
	},
	{
		// pair of square brackets
		"{}",
		1,
		1,
	},

}