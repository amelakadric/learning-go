package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	//Output info
	fmt.Println("BMI Calculator")
	fmt.Println("-----------------")

	//Prompt for user input (weight + height)
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')


	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')
	
	//it includes the \n from input
	//cleaning the input

	weightInput = strings.Replace(weightInput, "\r\n", "", -1)
	heightInput = strings.Replace(heightInput, "\r\n", "", -1)

	fmt.Print(weightInput)

	fmt.Print(heightInput)


	// Save that user input in variables

	//Calculate the BMI (weight / (height * height))

	//Output the calculated BMI
}