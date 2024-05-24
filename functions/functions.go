package main

import "fmt"

// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:
//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hello ", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func dumbFunc() string {
	return "Hola!"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer

func threeAdder(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func randomNumber() int {
	return 5
}

// * Write a function that returns any two numbers
func twoRandomNumber() (int, int) {

	return 1, 2
}

//	func twoRandomNumber () []int {
//		var result = []int {5,6}
//		return result
//	}
//
// * Add three numbers together using any combination of the existing functions.
//   - Print the result
func randomSum() {
	num1, num2 := twoRandomNumber()
	fmt.Println(randomNumber() + num1 + num2)

}
func main() {

	greetPerson("Sajib")
	fmt.Println("Sum of three numbers: ", threeAdder(1, 2, 3))
	randomSum()

	fmt.Println(dumbFunc())

	//* Call every function at least once

}
