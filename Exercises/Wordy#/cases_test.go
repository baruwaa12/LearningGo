package wordy

// Source: exercism/problem-specifications
// Commit: 51c9b0b wordy: Test question `What is?` because it's too short
// Problem Specifications Version: 1.5.0

type wordyTest struct {
	description string
	question    string
	ok          bool
	answer      int
}

var tests = []wordyTest{
	{
		"just a number",
		"What is 5 plus 7?",
		true,
		12,
	},
	{
		"just a number",
		"What is 5 plus 7",
		true,
		12,
	},
	{
		"addition",
		"What is 1 plus 1?",
		true,
		2,
	},
	{
		"more addition",
		"What is 53 plus 2?",
		true,
		55,
	},
	{
		"addition with negative numbers",
		"What is -1 plus -10?",
		true,
		-11,
	},
	{
		"large addition",
		"What is 123 plus 45678?",
		true,
		45801,
	},
	{
		"subtraction",
		"What is 4 minus -12?",
		true,
		16,
	},
	{
		"multiplication",
		"What is -3 multiplied by 25?",
		true,
		-75,
	},
	{
		"division",
		"What is 33 divided by -3?",
		true,
		-11,
	},	
}