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
		oldVerse: []string{"This is the Cat", "that lay in in the house that jack built"},
		newVerse: []string{"This is the BamBam", "that frightened the cat", "that lay in in the house that jack built"},

	},
}


var replaceThisIsWithThatAndActionTestCases  = []struct {
	line string
	action string
	newLine  string
}{
	{
		line: "This is the Cat",
		action: "frightened",
		newLine: "that frightened the cat",
	},
}

