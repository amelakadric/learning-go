package scopes

import "fmt"

var userName = "Amela"

func main() {
	shouldContinue := true
	userName := "Maximilian" //this shadows the higher-level variable
	// ovo ce da se printa, pravi se nova variabla,
	// stara se ne menja

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
