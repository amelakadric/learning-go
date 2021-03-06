package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/learning-go/02-BMI-Calculator-App/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics()(weight float64, height float64){

	weight = getUserInput(info.WeightPrompt)
	//it includes the \n from input
	//cleaning the input
	
	height = getUserInput(info.HeightPrompt)

	return 
}

func getUserInput(promptText string) (value float64){
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	value, _= strconv.ParseFloat(userInput, 64);
	return 
}