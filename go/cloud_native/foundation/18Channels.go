package main

import "fmt"

/**
Although main and the anonymous goroutine run concurrently and could in theory run in any order, the blocking behavior of unbuffered channels guarantees that the output will always be “ping” followed by “pong.”
*/

func main() {
	ch := make(chan string) // Allocate a string channel

	go func() {
		message := <-ch      // Blocking receive; assigns to message
		fmt.Println(message) // "ping"
		ch <- "pong"         // Blocking send
	}()

	ch <- "ping"      // Send "ping"
	fmt.Println(<-ch) // "pong"

	ch_buffer := make(chan string, 2)

	ch_buffer <- "foo"
	ch_buffer <- "bar"

	// Two non-blocking receive
	fmt.Println(<-ch_buffer)
	fmt.Println(<-ch_buffer)

	// Blocking receive
	//fmt.Println(<-ch_buffer)

}
