package domain

import "math"

type Point struct {
	x, y float64
}

func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}
func NewPoint(x, y float64) Point {
	return Point{x, y}
}
