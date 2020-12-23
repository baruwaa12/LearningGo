package aliquot

package perfect

var classificationTestCases = []struct {
	description string
	input       int64
	ok          bool
	expected    Classification
}{
	{
		description: "Smallest perfect number is classified correctly",
		input:       6,
		ok:          true,
		expected:    ClassificationPerfect,
	},
	{
		description: "Medium perfect number is classified correctly",
		input:       28,
		ok:          true,
		expected:    ClassificationPerfect,
	},
	{
		description: "Large perfect number is classified correctly",
		input:       33550336,
		ok:          true,
		expected:    ClassificationPerfect,
	},
	{
		description: "Large perfect number is classified correctly",
		input:       33550336,
		ok:          true,
		expected:    ClassificationPerfect,

	}