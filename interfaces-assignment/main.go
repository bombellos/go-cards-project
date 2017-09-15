package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{10}
	tr := triangle{5, 10}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println("Area of", fmt.Sprintf("%T", s), "is equal", s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}
