package triangle

import "math"

type Kind int 

const (
	NaT = iota
	Equ        // equilateral 
	Iso        // isosceles
	Sca        // scalene
)

// Floats type
type Floats []float64

func valid(v float64) bool {
	return v > 0 && !math.IsNaN(v) && !math.IsInf(v, 0)
}

func (floats Floats) forall(pred func(f float64) bool) bool {
	for _, float := range floats {
		if !pred(float) {
			return false
		}
	}
	return true
}

// KindFromSides functions
func KindFromSides(a, b, c float64) Kind {
	floats := Floats{a, b, c}
	if !floats.forall(valid) {
		return NaT
	}
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && a == c && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}