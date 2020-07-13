package pascaltriangle

var testCases = []struct {
	description string
	rows        int
	expected    [][]int
	}{
		{
			description: "Build a valid pascals triangle with 3 rows",
			rows:       3,
			expected:    [][]int{{1}, {1,1}, {1,1,1}},
		},
		{
			description: "Build a valid pascals triangle with 5 rows",
			rows:       5,
			expected:    [][]int{{1}, {1,1}, {1,1,1}, {1,1,1,1}, {1,1,1,1,1}},
		},
		
	}