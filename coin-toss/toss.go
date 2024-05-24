//--Summary:
//  Create a program that can perform Simulate the 1000 coin toss hypothesis.
//  If you coin toss and eliminate the persons who have tails, there will be one person who always gets head.

//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
)

func tossCoin() int {
	result := rand.Intn(2) //0 means head, 1 means tails
	return result
}

var numberOfToss = 0

func main() {
	fmt.Println("#################### S #################################")
	fmt.Println("#################### T #################################")
	fmt.Println("#################### A #################################")
	fmt.Println("#################### R #################################")
	fmt.Println("#################### T #################################")

	var results map[int]int
	results = make(map[int]int)

	for i := 0; i < 1000; i++ {
		results[i+1] = tossCoin()
	}

	for key, value := range results {
		if value != 0 {
			delete(results, key)
		}
	}
	fmt.Println("Contestant after first Round:", len(results))
	numberOfToss = 1
	tossTillEnd(results)
}

func tossTillEnd(result map[int]int) {

	fmt.Println("------------------------------------------------------------------------------------")
	numberOfToss += 1
	fmt.Println("Number of Toss:", numberOfToss)
	for key := range result {
		result[key] = tossCoin()

	}
	fmt.Println("Number of Contestant Before Eliminating: ", len(result))

	for key, value := range result {
		if value != 0 {
			delete(result, key)
		}
	}

	newLength := len(result)
	fmt.Println("Number of Contestant After Eliminating: ", newLength)

	if len(result) > 1 {

		tossTillEnd(result)
	} else if len(result) == 0 {
		fmt.Println("We Do not have any winner!")
	} else {

		fmt.Println("The winner is", result)
	}

}
