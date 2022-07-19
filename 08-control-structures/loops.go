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
		calculateSumUpToNumber()
	case "4":
		calculateListSum()
	}
}

func calculateSumUpToNumber() {
	// chosenNumber, err := getInputNumber()

	// if err!= nil {
	// 	fmt.Println("Invalid number input")
	// 	return
	// }

}
func calculateFactorial() {}

func calculateListSum() {}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid input")
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)
	return int(chosenNumber), nil
}

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice")
	fmt.Println("1) Add up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == "1" || userInput == "2" || userInput == "3" ||
		userInput == "4" {
		return userInput, err
	} else {
		return "", errors.New("INVALID INPUT")
	}

}
