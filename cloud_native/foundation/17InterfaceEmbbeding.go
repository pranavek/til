package main

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

//Third interface which links above interface
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	fmt.Println("vim-go")
}
