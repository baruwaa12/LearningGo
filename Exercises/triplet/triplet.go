package triplet


func tripletchecker(a int, b int, c int) int {
	var total int

	// if a < b < c
	// and if a + b +  c = 1000
	// and if a**2 + b**2 = c**2

	if (a < b && b < c) && (a*a + b*b == c*c){
		total = (c*c)
	}


	return total


}