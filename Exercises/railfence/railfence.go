package main

import (
	
	"strings"
)
 
func encode(s string, r int) string {
	b := board(len(s), r)
	sr := []rune(s)
	var builder strings.Builder
	for _, v := range b {
		builder.WriteRune(sr[v])
	
	}
	return builder.String()
	
}

func decoder( s string, r int) string {
	b := board(len(s), r)
	sr := []rune(s)
	decoded := make([]rune, len(s))
	for i, v := range b {
		decoded[v] = sr[i]
	}
}

func board(width, r int) (p []int) {
	var px, py = 0, 1
	pattern := make([][]int, r)
	for i := 0; i < width; i++ {
		pattern[px] = append(pattern[px], i)
		px += py
		if px == 0 || px == r-1 {
			py *= -1
		}
	}

	for _, row := range pattern {
		for _, b := range row {
			p = append(p, b) 

	}
	return p	
	}
}


