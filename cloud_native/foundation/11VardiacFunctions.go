package main

import "fmt"

func product(factors ...int) int {
	p := 1

	for _, n := range factors {
		p *= n
	}

	return p
}

func main() {
	fmt.Println(product(2, 2, 2))

	m := []int{3, 3, 3}
	fmt.Println(product(m...))
}
