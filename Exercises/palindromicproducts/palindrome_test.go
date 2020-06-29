package palindromic

var testCases = []struct {
	description string
	input       []int
	expected    []int
}{
	{
		description: "Given an array with one duplicate value return array with just unique values",
		input:       []int{1, 2, 3, 3},
		expected:    []int{1, 2, 3},
	},
	
	{
		description: "No duplicate values returns input array",
		input:       []int{1, 2, 3},
		expected:    []int{1, 2, 3},
	},
	
}
