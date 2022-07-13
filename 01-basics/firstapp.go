package main

import (
	"fmt"

	"github.com/amelakadric/learning-go/01-basics/greeting"
)


func main() {
	// var greetingText string
	// greetingText="Hello World!"

	// var greetingText string = "Hello world!"

	

	luckyNumber:= 17

	evenMoreLuckyNumber:= luckyNumber+5

	fmt.Println(greeting.GreetingText)
	fmt.Println(greeting.GreetingText)

	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(luckyNumber) / 3
	fmt.Println(newNumber)

	var defaultFloat float64 = 1.23456789123456789123456
	var smallFloat float32= 1.23456789123456789123456

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	var firstRune rune = '$'
	fmt.Println(string(firstRune))

	var firstByte byte = 'a'
	fmt.Println(string(firstByte))


	firstName :="Amela"
	lastName :="Kadric"
	fullName:= fmt.Sprintf("%v %v", firstName, lastName)
	age:=21

	fmt.Printf("Hi i am %v and i am %v (Type: %T) years old", fullName, age, age)

	

}