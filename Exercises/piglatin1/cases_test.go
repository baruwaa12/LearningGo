package piglatin

var testCases = []struct {
	description string
	word        string
	translated    string
	}{
		{
			description: "Vowel Sound",
			word:        "Alex",
			translated:  "Alexay",
		},
		{
			description: "Consonant",
			word:        "chair",
			translated:  "airchay",
		},
		
	}