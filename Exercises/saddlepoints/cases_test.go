package saddlepoints

var buildMatrixtestCases = []struct {
	description string
	column   int
	row		 int
	matrix   [][]int
	}{
		{
			description: "square - 2x2",
			column:        2,
			row:           2,
			matrix:     [][]int{ {}, {} },
		},
		{
			description: "square - 4x4",
			column:         4,
			row:            4,
			matrix:     [][]int{ {}, {}, {}, {} },
		},
		{
			description: "square - 3x3",
			column:         3,
			row:            3,
			matrix:     [][]int{ {}, {}, {} },
		},
	}