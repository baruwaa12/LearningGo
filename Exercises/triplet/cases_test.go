package triplet


var PythagoreanTriplets = []struct {
	description string
	a   int
	b	int
	c   int
	sum  int
	}{
		{
			description: "Most popular triplet",
			a: 3,
			b: 4,       
			c: 5,  
			sum: 12,
		},
		
		{
			description: "The answer to this question",
			a: 200,
			b: 375,       
			c: 425,  
			sum: 1000,
		},
		{
			description: "Doesnt obey the rules of pythagoras",
			a: 1,
			b: 2,       
			c: 1,  
			sum: 4,
		},
		
}
