package main

import "fmt"

func main() {

	age := 21

	fmt.Println(age)

	myAge := &age

	fmt.Println(*myAge)

}