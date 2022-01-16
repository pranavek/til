package main

import "fmt"

func incrementer() func() int {
	i := 0

	return func() int { // Returns an anonymous function
		i++
		return i //"closes over" parent function's i
	}
}
func main() {
	increment := incrementer()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	newIncrement := incrementer()
	fmt.Println(newIncrement())

}
