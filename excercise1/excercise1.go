package main

import "fmt"

func main() {
	firstName := "Amela"
	var lastName string = "Kadric"
	
	fmt.Println(firstName)
	fmt.Println(lastName)

	var currentYear int = 2022
	birthYear:=2001

	age:=currentYear-birthYear
	fmt.Println(age)


	currentYear++

	age=currentYear-birthYear
	fmt.Println(age)


}