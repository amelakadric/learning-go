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
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]

	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}

	prices = append(prices, discountPrices...)

	fmt.Println(prices)

}

// func main() {
// var productNames [4]string = [4]string{"A Book"}
// prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// fmt.Println(prices)

// productNames[2] = "A Carpet"

// fmt.Println(productNames)

// fmt.Println(prices[3])

// // featuredPrices := prices[1:3] //starts with element 1 excluding element 3

// featuredPrices := prices[1:] //starts at the beginning
// // featuredPrices := prices[1:]//ends on the end

// featuredPrices[0] = 199.99
// highlightedPrices := featuredPrices[:1]

// fmt.Println(highlightedPrices)
// fmt.Println(prices)
// fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// highlightedPrices = highlightedPrices[:3]
// fmt.Println(highlightedPrices)
// fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// //editing slices is editing the array also
// //beacuse its just like a window into an array
//
//https://www.youtube.com/watch?v=j7eEt_jk1V0
// F U
// }
