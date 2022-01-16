package main

import (
	"fmt"
	"os"
)

func ifs() {
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// Initilization precedes the condition
	if _, err := os.Open("foo.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All is fine")
	}

}

func main() {
	ifs()
}
