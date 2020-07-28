package rna

var testCases = []struct {
	description string
	dna string
	result string
}{
	{
		description: "Random DNA string",
		dna: "ACGTGGTCTTAA",
		result: "UGCACCAGAAUU",
	},
}