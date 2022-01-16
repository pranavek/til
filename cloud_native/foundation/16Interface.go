package main

import (
	"fmt"
	"math"
)

/*
* - Describe general behaviour without being coupled to implementation details
* - An interface can thus be viewed as a contract that a type may satisfy, opening door to powerful abstraction
 */

type Shape interface {
	Area() float64
}

//Let's define struct which is a shape
type Circle struct {
	Radius float64
}

type Rectanagle struct {
	Width, Height float64
}

// Let's define method for the shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectanagle) Area() float64 {
	return r.Height * r.Width
}

// Let's create a function to print the area
func PrintArea(s Shape) {
	fmt.Printf("%T's area is %0.2f\n", s, s.Area())
}

func main() {
	r := Rectanagle{Width: 10, Height: 10}
	PrintArea(r)

	c := Circle{Radius: 10}
	PrintArea(c)

	//Assertions
	var s Shape
	s = Circle{}       // s is an expression of Shape
	circ := s.(Circle) // Assert that s is a Circle
	fmt.Printf("%T\n", circ)

}
