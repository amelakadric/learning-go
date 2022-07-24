package main

import "fmt"

func main() {
	// new & make

	// hobbies := []string{"Sports", "Reading"}
	// composite literal making of a slice

	hobbies := make([]string, 2, 10)

	hobbies[0] = "Sports"
	hobbies[1] = "Reading"

	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)

}
