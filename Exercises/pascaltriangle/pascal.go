package pascaltriangle

import (
	"fmt"
)

func createDefaultTriangle(rows int) (triangle [][]int) {
	triangle = make([][]int, rows)

	for i := 0; i < rows; i++ {
		te := make([]int, i+1)
		for r := 0; r < len(te); r++ {
			te[r] = 1
		}
		triangle[i] = te

	}
	fmt.Println(triangle)
	return triangle
}
