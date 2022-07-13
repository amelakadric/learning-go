package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/learning-go/02-BMI-Calculator-App/info"
)



func main() {

	//Output info
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)

	//Prompt for user input (weight + height)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')


	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')
	
	//it includes the \n from input
	//cleaning the input

	// Save that user input in variables

	weightInput = strings.Replace(weightInput, "\r\n", "", -1)
	heightInput = strings.Replace(heightInput, "\r\n", "", -1)

	weight, _:= strconv.ParseFloat(weightInput, 64);
	height, _:= strconv.ParseFloat(heightInput, 64);

	//Calculate the BMI (weight / (height * height))
	bmi:= weight/(height*height)


	//Output the calculated BMI
	fmt.Printf("Your BMI: %.2f", bmi)
}