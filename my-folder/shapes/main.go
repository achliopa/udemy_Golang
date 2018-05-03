package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	triangle1 := triangle{
		height: 5.5,
		base:   3.2,
	}

	square1 := square{
		sideLength: 3.6,
	}

	printArea(triangle1)
	printArea(square1)

}

func printArea(s shape) {
	fmt.Println("Shape area: ", s.getArea())
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
