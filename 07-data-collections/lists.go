package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[3])

	// featuredPrices := prices[1:3] //starts with element 1 excluding element 3

	featuredPrices := prices[:3] //starts at the beginning
	// featuredPrices := prices[1:]//ends on the end

	highlightedPrices := featuredPrices[:1]

	fmt.Println(highlightedPrices)

	//https://www.youtube.com/watch?v=j7eEt_jk1V0
	// F U
}
