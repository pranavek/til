package main

import (
	"fmt"
	"time"
)

func switchs() {
	i := 0

	switch i % 3 {
	case 0:
		fmt.Println("zero")
		fallthrough
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("I don't know")
	}

	switch hour := time.Now().Hour(); { // Empty expression means "true"
	case hour >= 5 && hour < 9:
		fmt.Println("I'm writing")
	case hour >= 9 && hour < 18:
		fmt.Println("I'm working")
	default:
		fmt.Println("I'm sleeping")
	}
}
func main() {
	switchs()
}
