package house1


var testCases = []struct {
	noun string
	action string
	oldVerse []string
	newVerse [] string

}{
	{
		noun: "BamBam",
		action: "frightened",
		oldVerse: []string{"this is the cat", "that lay in in the house that jack built"},
		newVerse: []string{"this is the BamBam", "that frightened the cat", "that lay in in the house that jack built"},

	},
}


var replaceThisIsWithThatAndActionTestCases  = []struct {
	line string
	action string
	newLine  string
}{
	{
		line: "this is the cat",
		action: "frightened",
		newLine: "that frightened the cat",
	},
}

var newLineTestCases  = []struct {
	noun string
	newLine  string
}{
	{
		noun: "cat",
		newLine: "this is the cat",
	},
}



