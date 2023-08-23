package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func printArea(s shape) {
	fmt.Println("The area is as such:", s.getArea())
}

func main() {
	fmt.Println("Hello World, Here is how this works")
	square := square{sideLength: 2}
	triangle := triangle{base: 2, height: 2}
	printArea(square)
	printArea(triangle)
}
