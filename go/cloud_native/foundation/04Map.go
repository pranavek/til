package main

import "fmt"

func maps() {
	freezing := make(map[string]float32)

	freezing["celsius"] = 0.0
	freezing["fahrenheit"] = 32.0
	freezing["kelvin"] = 273.2

	fmt.Println(freezing["kelvin"])
	fmt.Println(len(freezing))

	delete(freezing, "kelvin")
	fmt.Println(len(freezing))

	test := map[string]int{
		"c": 0,
		"k": 1,
		"f": 2,
	}

	fmt.Println(test["f"])

	//Map membership testing
	test_val := test["non_exist_key"]
	fmt.Println(test_val)

	test_val2, ok := test["non_exist_key"]
	fmt.Println(test_val2)
	fmt.Println(ok)
}

func main() {
	maps()
}
