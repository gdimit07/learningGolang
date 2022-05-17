package main

import "fmt"

type shape interface {
	printArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func (t triangle) printArea() float64 {

	return t.height * t.base
}

func (s square) printArea() float64 {
	return s.sideLength * s.sideLength
}

func getArea(s shape) float64 {
	return s.printArea()
}

// func getArea(t triangle) float64 {
// 	return 0.5 * t.base * t.height
// }

func main() {

	tri := triangle{height: 3.1, base: 2.1}
	squ := square{sideLength: 2}

	fmt.Println("triangle area:", getArea(tri))
	fmt.Println("square area:", getArea(squ))
}
