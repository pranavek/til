package main

import "fmt"

/**
* Append function returns the appended slice rather than modifiying the
* existing slice
 */
func slices() {
	n := make([]int, 3)

	fmt.Println(n)
	fmt.Println(len(n))

	n[0] = 10

	fmt.Println(n[0])

	//Slice literal
	m := []int{1} //In comparision with the array, the size is not specified
	m = append(m, 2)

	fmt.Println(m)

	//Append a slice to itself
	m = append(m, m...)
	fmt.Println(m)

	s1 := m[:2]
	fmt.Println(s1)
	s2 := m[2:]
	fmt.Println(s2)

	//Change propogation
	m[1] = 20
	fmt.Println(m)
	fmt.Println(s1)

	//String
	s := "foö"
	r := []rune(s)
	b := []byte(s)

	fmt.Printf("%7T %v\n", s, s)
	fmt.Printf("%7T %v\n", r, r)
	fmt.Printf("%7T %v\n", b, b)

}

func main() {
	slices()
}
