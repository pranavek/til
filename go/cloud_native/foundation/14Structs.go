package main

import "fmt"

type Cordinates struct {
	lat int
	lon int
}

func main() {
	var cor Cordinates
	fmt.Println(cor) // Struct is never nil

	cor = Cordinates{} // Define a emtpy sturct
	fmt.Println(cor)

	cor = Cordinates{1, 2}
	fmt.Println(cor)

	cor = Cordinates{lat: 2}
	fmt.Println(cor)

	cor.lon = 20

	fmt.Println(cor)

	//Structs are commonly manipulated by using pointers
	var c *Cordinates = &Cordinates{1, 2}
	fmt.Println(c)

	c.lat, c.lon = c.lon, c.lat
	fmt.Println(c)

}
