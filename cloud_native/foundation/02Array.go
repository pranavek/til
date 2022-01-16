package main

import "fmt"

func array() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(a[1])

	b := [3]int{2, 4, 6}
	c := [...]int{2, 4, 6, 8}

	fmt.Println(len(c))
	fmt.Println(b[len(b)-1])
}

func main() {
	array()
}
