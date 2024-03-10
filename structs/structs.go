//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func calculateArea(length, width int) int {
	return length * width
}

func calculatePermiter(length, width int) int {
	return 2 * (length + width)
}

func main() {

	firstRect := Rectangle{10, 8}
	fmt.Println(firstRect)
	fmt.Println(calculateArea(firstRect.length, firstRect.width))
	firstRect.width = firstRect.width * 2
	firstRect.length = firstRect.length * 2
	fmt.Println(calculateArea(firstRect.length, firstRect.width))
	fmt.Println(calculatePermiter(firstRect.length, firstRect.width))

}
