package queenattack

import (
    "errors"
)


func filterInput(w, b string) error {
	// check input length
	if len(w) != 2 || len(b) != 2 {
		return errors.New("wrong input length")
	}
	// check if they are on the same square
	if w == b {
		return errors.New("same square")
	}
	// check letter axis limits
	if w[0] > 'h' || w[0] < 'a' || b[0] > 'h' || b[0] < 'a' {
		return errors.New("out of range letters axis")
	}
	// check number axis limits
	if w[1] > '8' || w[1] < '1' || b[1] > '8' || b[1] < '1' {
		return errors.New("out of range number axis")
	}
    return nil
}

func abs(n int) int {
    if n == 0 {
        return 0
    }
    if n < 0 {
        return -n
    }
    return n
}


// CanQueenAttack checks if the queens can attack and if their position is valid
func CanQueenAttack(w, b string) (bool, error) {
	if err := filterInput(w, b); err != nil {
		return false, err
	}
	// check same row or same column
	if w[0] == b[0] || w[1] == b[1] {
		return true, nil
	}
	// check diagonal positive and negative
	if abs(int(w[0])-int(b[0])) == abs(int(w[1])-int(b[1])) {
		return true, nil
	}
	return false, nil

}