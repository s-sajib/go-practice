//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type product struct {
	price int
	name  string
}

func printStats(shoppingList []product) {

	totalPrice := 0
	totalNumberOfItems := 0

	for i := 0; i < len(shoppingList); i++ {
		itemName := shoppingList[i].name
		itemPrice := shoppingList[i].price

		if itemName != "" {
			totalPrice += itemPrice
			totalNumberOfItems = totalNumberOfItems + 1
		}
	}

	fmt.Println("Last product on list: ", shoppingList[totalNumberOfItems-1])
	fmt.Println("Total number of items on list: ", totalNumberOfItems)
	fmt.Println("Total Cost of items: ", totalPrice)
}
func main() {

	shoppingList := [4]product{{10, "apple"}, {20, "orange"}, {30, "banana"}}
	printStats(shoppingList[:])

	shoppingList[3] = product{100, "Biriyani"}
	printStats(shoppingList[:])
}
