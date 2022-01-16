package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)

	select {
	case <-ch1:
		fmt.Println("test")
	case x := <-ch2:
		fmt.Println(x)
	case ch3 <- "hi":
		fmt.Println("Triggered ch3")
	default:
		fmt.Println("I didn't get anything")
	}

}
