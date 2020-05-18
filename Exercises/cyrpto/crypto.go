package crypto

import (
	"math"
	"regexp"
	"strings"
)

var textRe = regexp.MustCompile("[^a-z0-9]+")

// Return columns and rows

func dim(numChars int) (int, int) {
	cols := int(math.Ceil(math.Sqrt(float64(numChars))))
	if cols*(cols-1) >= numChars {
		return cols, cols - 1
	}
	return cols, cols
}

func chunk(bytes []byte, n int) [][]byte {
	var ret [][]byte
	for from, to := 0, n; to <= len(bytes); from, to = from+n, to+n {
		ret = append(ret, bytes[from:to])
	}
	return ret
}

func transpose(bytes [][]byte) [][]byte {
	rows := len(bytes)
	cols := len(bytes[0])
	ret := make([][]byte, cols)
	for c := 0; c < cols; c++ {
		ret[c] = make([]byte, rows)
		for r := 0; r < rows; r++ {
			ret[c][r] = bytes[r][c]
		}
	}
	return ret
}

// Encode Function
func Encode(text string) string {
	if text == "" {
		return ""
	}
	normalized := textRe.ReplaceAllString(strings.ToLower(text), "")
	l := len(normalized)
	cols, rows := dim(l)
	bytes := []byte(normalized + strings.Repeat(" ", cols*rows-l))
	var ss []string
	for _, row := range transpose(chunk(bytes, cols)) {
		ss = append(ss, string(row))
	}
	return strings.Join(ss, " ")
}

