package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int

	//ch = make(chan int, 1)
	//ch <- 1

	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(5 * time.Second):
		fmt.Println("Timed out")
	}
}
