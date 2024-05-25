//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Member struct {
	name string
}

type Book struct {
	name      string
	available bool
}

type Library struct {
	id           int
	bookId       int
	memberId     int
	lentTime     time.Time
	returnedTime time.Time
}
type Members map[int]Member
type Books map[int]Book

func checkOut(id int, bookId int, memberId int) Library {
	record := Library{
		id:       id,
		bookId:   bookId,
		memberId: memberId,
		lentTime: time.Now(),
	}

	fmt.Println("Record: ", record.lentTime)
	return record
}
func checkIn(returnTime *time.Time) {
	*returnTime = time.Now()

}

func main() {

	members := Members{1: {name: "Sajib"}, 2: {name: "Tisha"}, 3: {name: "Mou"}, 4: {name: "Surad"}}

	fmt.Println("Members: ", members)
	books := Books{1: {name: "Bangla", available: true}, 2: {name: "English", available: true}, 3: {name: "Math", available: true}, 4: {name: "Physics", available: true}}
	fmt.Println("Books: ", books)
	records := []Library{checkOut(1, 1, 1)}
	fmt.Println("Checked Out: ", records)
	checkIn(&records[0].returnedTime)
	fmt.Println("Checked In: ", records)
}
