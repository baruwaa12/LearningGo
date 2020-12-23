package Alphametics

// Solve the alphametics puzzle
var testCases = []struct {
	description   string
	input         string
	expected      map[string]int
	errorExpected bool
}{
	{
		description: "puzzle with three letters",
		input:       "I + BB == ILL",
		expected:    map[string]int{"B": 9, "I": 1, "L": 0},
	},
	{
		description:   "solution must have unique value for each letter",
		input:         "A == B",
		errorExpected: true,
	},
	{
		description:   "leading zero solution is invalid",
		input:         "ACA + DD == BD",
		errorExpected: true,
	},
	{
		description: "puzzle with four letters",
		input:       "AS + A == MOM",
		expected:    map[string]int{"A": 9, "M": 1, "O": 0, "S": 2},
	},
}