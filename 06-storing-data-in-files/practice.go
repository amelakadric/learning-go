package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Product struct {
	id          string
	name        string
	description string
	price       float64
}

func (product *Product) outputData() {
	fmt.Printf("Name: %v,\n id: %v,\n Description: %v, \n Price: %.2f$\n\n", product.name, product.id, product.description, product.price)
}

func NewProduct(id string, name string, desc string, price float64) *Product {

	return &Product{id: id, name: name, description: desc, price: price}
}

func main() {

	product1 := Product{"123", "Alchemist", "Awesome book", 9.99}

	product2 := *NewProduct("456", "Juice", "Refreshing beverage", 2.49)

	product3 := *getProduct()

	// fmt.Println(product1)
	// fmt.Println(product2)

	product1.outputData()
	product2.outputData()
	product3.outputData()
}

func getProduct() *Product {
	var prod *Product
	prod.id = readUserInput("Id of product: ")
	prod.name = readUserInput("Name of product: ")
	prod.description = readUserInput("Short description: ")
	prod.price, _ = strconv.ParseFloat(readUserInput("Price of product: "), 64)
	return prod
}

func readUserInput(promptText string) string {
	fmt.Print(promptText)

	data, _ := reader.ReadString('\n')

	data = strings.Replace(data, "\r\n", "", -1)

	return data

}

// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.
