package main

import "fmt"

type Vertex struct {
	x, y int
}

/*
* - Methods are functions attached to the type
* - The syntax is similar to that of a function, expect that it includes
* extra receiver argument before the function name
* - It specifies the type that the method is attached to
* - When the method is called, the instance is accssible by the name specified in the receiver
 */

//Extend the capabilites of Vertex sturct using Methods
func (v *Vertex) square() {
	v.x *= v.x
	v.y *= v.y
}

type MyMap map[int]int

func (m MyMap) Length() int {
	return len(m)
}

func main() {
	var v *Vertex = &Vertex{x: 2, y: 3}
	fmt.Println(v)

	v.square()
	fmt.Println(v)

	// Extend standard type using type keyword
	m := MyMap{1: 1, 2: 2}
	fmt.Println(m.Length())
}
