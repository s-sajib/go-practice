//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
)

func rollDice() int {
	return rand.Intn(6) + 1
}

func main() {
	numberOfDice := 2

	var sum int = 0

	for i := 0; i < numberOfDice; i++ {
		sum = sum + rollDice()
	}

	switch {
	case sum == 2 && numberOfDice == 2:
		fmt.Println("Snake Eyes!")
	case sum == 7:
		fmt.Println("Lucky 7")
	case sum%2 == 0:
		fmt.Println("Even")
	case sum%2 != 0:
		fmt.Println("Odd")
	}

}
