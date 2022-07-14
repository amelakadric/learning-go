package main

import (
	"fmt"
)

func main() {

	age := 21

	fmt.Println(age)

	// myAge := &age

	var myAge *int
	myAge = &age

	fmt.Println(*myAge)

	*myAge = 22

	fmt.Println(*myAge)
	fmt.Println(age)

	doubledAge := double(myAge)
	fmt.Println(doubledAge)
	
	double(myAge)

	fmt.Println(age)
}


func double(number *int) int {
	result := *number * 2 // it doesnt make uneccessarry copies of values
	*number=100
	return result
}