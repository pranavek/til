package main

import "fmt"

func logical() {
	and := true && false
	fmt.Println(and)

	or := true || false
	fmt.Println(or)

	not := !true
	fmt.Println(not)
}

func main() {
	logical()
}
