package main

import (
	"fmt"
	"image/color"
	"sync"
)

// Point Structure
type Point struct{ X, Y float64 }

//ColoredPoint Structure
type ColoredPoint struct {
	Point
	Color color.RGBA
		func ScaleBy struct {

		}
}

func colour() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

}

// Distance Function
func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

// Lookup Function
func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}
