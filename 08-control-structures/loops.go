package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice, exiting")
		return
	}

	switch selectedChoice {
	case "1":
		calculateSumUpToNumber()

	case "2":
		calculateFactorial()
	case "3":
		calculateSumManually()
	case "4":
		calculateListSum()
	}
}

func calculateSumManually() {
	isEnteringNumbers := true
	sum := 0
	fmt.Println("Keep on entering numbers, the program will quit when you enter any other value")

	for isEnteringNumbers {
		chosenNumber, err := getInputNumber()
		sum += chosenNumber

		isEnteringNumbers = (err == nil)

	}
	fmt.Printf("Result: %v", sum)
}

func calculateSumUpToNumber() {
	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid number input")
		return
	}

	var sum int

	for i := 1; i <= chosenNumber; i++ {
		sum += i
	}

	fmt.Printf("Result: %v\n", sum)

}
func calculateFactorial() {

	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()

	if err != nil {
		fmt.Println("Invalid number input")
		return
	}
	fact := 1

	for i := 1; i <= chosenNumber; i++ {
		fact *= i
	}

	fmt.Println("Result: %v\n", fact)

}

func calculateListSum() {
	fmt.Println("Please enter a comma-separated list of numbers.")
	inputNumberList, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	inputNumberList = strings.Replace(inputNumberList, "\r\n", "", -1)
	inputNumbers := strings.Split(inputNumberList, ",")

	sum := 0

	for index, value := range inputNumbers {
		fmt.Printf("index: %v, Value: %v\n", index, value)
		number, err := strconv.ParseInt(value, 0, 64)

		if err != nil {
			continue
		}
		sum += int(number)
	}

	fmt.Printf("Result: %v\n", sum)

}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\r\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)
	return int(chosenNumber), err
}

func getUserChoice() (string, error) {
	fmt.Println("1) Add up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers")
	fmt.Println("Please make your choice: ")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\r\n", "", -1)

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, err
	} else {
		return "", errors.New("INVALID INPUT")
	}

}
