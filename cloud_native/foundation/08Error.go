package main

import "fmt"

type NestedError struct {
	Message string
	Err     error
}

func (e *NestedError) Error() string {
	return fmt.Sprintf("%s\n contains: %s", e.Message, e.Err.Error())
}

func errors() {
	e1 := errors.New("Pranav's error")
	e2 := fmt.Errorf("Pranav's %d error", 1)

	// Error is an interface, you can create your own interface

}

func main() {
	errors()
}
