package perfectnumber

var testCases = []struct {
	description string
	number       int
	types  	 string
	}{
		{
			description: "Smallest Perfect Number",
			number:        6,
			types:     "perfect",
		},
		{
			description: "Smallest Abundant Number",
			number:         16,
			types:     "abundant",
		},
		{
			description: "Deficient Number",
			number:         8,
			types:     "deficient",
		},

		
	}