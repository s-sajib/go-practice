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

func main() {
	type product struct {
		price int
		name  string
	}

	shoppingList := [4]product{{10, "apple"}, {20, "orange"}, {30, "banana"}}
	fmt.Println(shoppingList[2])
	length := len(shoppingList)
	fmt.Println(length)

	sum := 0
	for i := 0; i < length; i++ {
		sum = sum + shoppingList[i].price
	}
	fmt.Println(sum)

	shoppingList[3] = product{40, "jambura"}
	fmt.Println(shoppingList)
}
