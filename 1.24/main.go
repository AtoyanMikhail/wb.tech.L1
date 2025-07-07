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

func Distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(a.x - b.x), 2) + math.Pow(math.Abs(a.y - b.y), 2))
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(1, 1)
	dist := Distance(*a, *b)
	fmt.Printf("Distance: %.4f", dist)
}