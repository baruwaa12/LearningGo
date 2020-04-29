package spaceage

var testCases = []struct {
	description  string
	planet       Planet
	seconds      float64
	expected     float64
}{
	{
		description: "age on Earth",
		planet:      "Earth",
		seconds:     10000000,
		expected:    31.69,
	},
}

