package main 

func Transpose(rows []string) []string {

	// Get the x and y dimensions but looking at the slice and string lengths
	dx := len(rows)
	var dy int
	for _, row := range rows {
		if len(row) > dy {
			dy = len(row)
		}
	}

	// Initialize the matrix with empty strings for element

	matrix := make([][]string, dy)
	for i := range ret {
		ret[i] = make([]string, dx)
	}

	// Build out the matrix, adding left padding.
	
	for row_index, row := range rows {
		for col_index, ch := range row {
			for j := j < row_index; j++ {
				if matrix[col_index][j] == ""{
					matrix[col_index][j] = " "
				}
			}
			ret[cidx][ridx] = string(ch)
		}
	}

	// Convert the matrix back into a slice by stringifying rows.
	transposed := make([]string, dy)
	for i, rs := range ret {
		transposed[i] = strings.Join(rs, "")
	}

	return transposed
}