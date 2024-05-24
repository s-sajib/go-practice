//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

// * Create a function to print out the contents of the assembly line
func printAsselbly(parts []Part) {
	for i := 0; i < len(parts); i++ {
		fmt.Println(parts[i])
	}
}

func main() {
	// * Using a slice, create an assembly line that contains type Part
	// * Perform the following:
	//   - Create an assembly line having any three parts
	assembly := []Part{"Part A", "Part B", "Part C"}
	printAsselbly(assembly)
	fmt.Println("---------------------------------")
	//   - Add two new parts to the line
	assembly = append(assembly, "Part D", "Part E")
	printAsselbly(assembly)

	fmt.Println("---------------------------------")
	//   - Slice the assembly line so it contains only the two new parts
	shortAssembly := assembly[3:]
	printAsselbly(shortAssembly)

}
