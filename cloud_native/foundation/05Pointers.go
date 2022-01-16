package main

import "fmt"

/**
* Stores the address of a variable
* Access a pointer = &
* Dereference a pointer = *
* p = &a ; The variable p points to a with it's type as *int - The * indicate it's a pointer type
 */
func pointers() {

	var a int = 10
	var p *int = &a

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 20
	fmt.Println(a)
}

func main() {
	pointers()
}
