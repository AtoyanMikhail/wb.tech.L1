package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p Point) Distance(another Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(p.x - another.x), 2) + math.Pow(math.Abs(p.y - another.y), 2)) 
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(1, 1)
	dist := a.Distance(*b)
	fmt.Printf("Distance: %.4f\n", dist)
}