package main

import "fmt"

func variables() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("int %d\n", i)
	fmt.Printf("float %f\n", f)
	fmt.Printf("bool %t\n", b)
	fmt.Printf("string %s\n", s)

}

func main() {
	variables()
}
