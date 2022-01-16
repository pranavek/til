package main

import "fmt"

func zeroByvalue(x int) {
	x = 0
}

func zeroByReference(x *int) {
	*x = 0
}

/**
* Refrences to the memory location
* Changs made to such reference types in function can affect the caller
* without needing to explicitly dereferencing them
 */

func update(x map[string]int) {
	x["x"] = 10
}

func main() {
	x := 5

	zeroByvalue(x)
	fmt.Println(x)

	zeroByReference(&x)
	fmt.Println(x)

	m := map[string]int{"a": 0, "b": 1}

	fmt.Println(m)
	update(m)
	fmt.Println(m)
}
