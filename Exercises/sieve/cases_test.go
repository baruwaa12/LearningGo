package sieve


var testCases = []struct {
	description string
	limit int
	list  []int
}{
	{
		description: "first prime number",
		limit: 2,
		list: []int{ 2 },
	},
	{
		description: "prime numbers upto 7",
		limit: 7,
		list: []int{ 2, 3, 5, 7},
	},
	{
		description: "prime numbers up to 29",
		limit: 29,
		list: []int{ 2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
	},
	{
		description: "prime numbers up to 89",
		limit: 89,
		list: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89,},
	},

}