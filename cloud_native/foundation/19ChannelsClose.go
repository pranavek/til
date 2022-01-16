package main

import "fmt"

func main() {
	ch := make(chan string, 10)

	ch <- "foo"

	close(ch)

	msg, ok := <-ch
	fmt.Printf("%q, %v\n", msg, ok)

	msg, ok = <-ch
	fmt.Printf("%q, %v\n", msg, ok)

}
