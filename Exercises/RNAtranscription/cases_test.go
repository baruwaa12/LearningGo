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
	{
		description: "Random DNA string 2",
		dna: "ACGTGGTCTTAA",
		result: "UGCACCAGAAUU",
	},
	{
		description: "Random DNA string 3",
		dna: "ACGCGCCTTAA",
		result: "UGCGCGGAAUU",
	},
	{
		description: "Random DNA string 4",
		dna: "CGCGTTTTAA",
		result: "GCGCAAAAUU",
	},

}