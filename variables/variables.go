//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favouriteColor string = "skyBlue"
	var birthYear, age int = 1998,25
	var (
		first string = "S"
		last string = "S"
	)
	var myAgeInDays int = 365 * age

	fmt.Printf("favourite color: %s \n", favouriteColor)
	fmt.Printf("Birth year: %d , Age: %d \n", birthYear, age)
	fmt.Printf("First: %s , Last: %s \n", first, last)
	fmt.Printf("Age in Days: %d \n", myAgeInDays)
}
