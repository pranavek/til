package main

import "fmt"

func fors() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	//for i< 10 {..}

	//Looping over arrays and slices
	s := []int{1, 2, 3, 4, 5, 6}

	for i, v := range s {
		fmt.Println(i, "->", v)
	}

	//Looping over maps
	month := map[int]string{
		1: "Jan",
		2: "Feb",
		3: "Mar",
	}

	for k, v := range month {
		fmt.Println(k, "->", v)
	}

}

func main() {
	fors()
}
