package main

import "fmt"

var userName = "Amela"

func main() {
	shouldContinue := true
	userName := "Maximilian"

	if shouldContinue {
		userAge := 32

		fmt.Printf("Name: %v, Age: %v", userName, userAge)
	}

	// fmt.Println(userAge) scope

}

func prinData() {
	fmt.Println(userName)
	// fmt.Println(shouldContinue) doesnt work
}
