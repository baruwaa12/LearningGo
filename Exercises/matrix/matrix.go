package matrix

import (
    "errors"
    "strings"
    "strconv"
)

// Matrix is a 2D matrix
type Matrix [][]int

// Rows returns the rows of a 2D matrix
func (m Matrix) Rows() (rows [][]int) {
    for i := 0; i < len(m); i++ {
        var row []int
        for j := 0; j < len(m[i]); j++ {
            row = append(row, m[i][j])
        }
        rows = append(rows, row)
    }
    return rows
}

// Cols returns the columns of a 2D matrix
func (m Matrix) Cols() (cols [][]int) {
    ncols := len(m[0])
    for j := 0; j < ncols; j++ {
        var col []int
        for i := 0; i < len(m); i++ {
            col = append(col, m[i][j])
        }
        cols = append(cols, col)
    }
    return cols
}

// Set sets the value of a matrix element
func (m Matrix) Set(r int, c int, v int) bool {
    if r < 0 || c < 0 { return false }
    if len(m.Rows()) - 1 < r { return false }
    if len(m.Cols()) - 1 < c { return false }
    m[r][c] = v
    return true
}

// New accepts a string and returns a pointer to a matrix
func New(s string) (*Matrix, error) {
    m := new(Matrix)
    rows := strings.Split(s, "\n")
    ncols := 0

    for _, row := range(rows){
        row = strings.TrimSpace(row)
        col := strings.Split(row, " ")
        if ncols == 0 { ncols = len(col) }
        if ncols != len(col) {
            return nil, errors.New("inconsistent number of columns")
        }
        var arr []int
        for _, val := range(col){
            v, err := strconv.Atoi(val)
            if err != nil {
                return nil, errors.New("parse error")
            }
            arr = append(arr, v)
        }
        *m = append(*m, arr)
    }

    return m, nil
}