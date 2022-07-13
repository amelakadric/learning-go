package main

import "fmt"

func main() {

	//the value of const has to be known before compilation
	const pi = 3.14
	radius := 5

	circumference := 2 * pi * float32(radius)

	fmt.Println(circumference)


	fmt.Printf("For a radius of %v, the circle circumference is %.2f", radius, circumference)

	
}