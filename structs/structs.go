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

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1 int
	x2 int
	y1 int
	y2 int
}

func calculateWidth(rect Rectangle) int {
	return int(math.Abs(float64(rect.x2 - rect.x1)))
}
func calculateLength(rect Rectangle) int {
	return int(math.Abs(float64(rect.y2 - rect.y1)))
}

func calculateArea(rectangle Rectangle) int {
	length := calculateLength(rectangle)
	width := calculateWidth(rectangle)
	return length * width
}

func calculatePermiter(rectangle Rectangle) int {
	length := calculateLength(rectangle)
	width := calculateWidth(rectangle)
	return 2 * (length + width)
}

func main() {

	firstRect := Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}
	fmt.Println(firstRect)
	fmt.Println(calculateArea(firstRect))
	firstRect.x2 = firstRect.x2 * 2
	firstRect.y2 = firstRect.y2 * 2
	fmt.Println(calculateArea(firstRect))
	fmt.Println(calculatePermiter(firstRect))

}
