package main

import (
	"fmt"
)

func main() {
	// new & make

	// hobbies := []string{"Sports", "Reading"}
	// composite literal making of a slice

	number := new(int)
	fmt.Println(number)
	fmt.Println(*number)

	anotherNumber := 0
	numberAdress := &anotherNumber
	fmt.Println(numberAdress)
	fmt.Println(anotherNumber)

	hobbies := make([]string, 2, 10)
	moreHobbies := new([]string) //returns a pointer
	// evenMoreHobbies:= []string{} // make([]string)

	*moreHobbies = append(*moreHobbies, "Swimming")
	fmt.Println(*moreHobbies)

	// aMap := make(map[string]int, 5)

	hobbies[0] = "Sports"
	hobbies[1] = "Reading"

	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)

}
