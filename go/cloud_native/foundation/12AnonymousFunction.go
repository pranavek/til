package main

import "fmt"

func sum(x, y int) int     { return x + y }
func product(x, y int) int { return x * y }

func main() {
	var f func(int, int) int // Function variable

	f = sum
	fmt.Println(f(2, 2))

	f = product
	fmt.Println(f(3, 3))

}
