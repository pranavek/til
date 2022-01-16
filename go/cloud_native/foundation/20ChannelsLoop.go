package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	ch <- "foo" // Send three (buffered) values to the channel
	ch <- "bar"
	ch <- "baz"

	close(ch) // Try without close, you will see what happens

	for s := range ch {
		fmt.Println(s)
	}
}
