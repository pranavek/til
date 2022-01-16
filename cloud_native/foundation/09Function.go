package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func add_three(x, y, z int) int {
	return x + y + z
}

func swap(x, y int) (int, int) {
	return y, x
}

func fact(x int) int {
	if x < 1 {
		return 1
	}

	return x * fact(x-1)
}

func defer_test() {
	defer fmt.Println("testing defer")
	fmt.Println("testing")
}

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(swap(10, 20))
	a, b := swap(20, 30)
	fmt.Println(a, "<->", b)
	fmt.Println(fact(3))
	defer_test()
}
