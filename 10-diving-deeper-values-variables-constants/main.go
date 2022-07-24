package main

import "fmt"

type person struct {
	name string
	age  int
}

type customNumber int

type personData map[string]person

func (num customNumber) pow(power int) customNumber {
	result := num
	for i := 1; i < power; i++ {
		result *= num
	}

	return result
}

func main() {

	var people personData = personData{
		"p1": {"Amela", 32},
	}

	fmt.Println(people)

	var dummyNumber customNumber = 5
	poweredNumber := dummyNumber.pow(3)

	fmt.Println(poweredNumber)

}
