package main

func (p *Point) ScaleBY(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type P *int
func (P) f() { /*...*/}