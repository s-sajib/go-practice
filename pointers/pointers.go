//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

package main

import "fmt"

var (
	Active   = true
	Inavtive = false
)

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
type SecurityTag struct {
	name string
	tag  bool
}

//* Create functions to activate and deactivate security tags using pointers

func activateTag(tag *bool) {
	*tag = Active
}
func deActivateTag(tag *bool) {
	*tag = Inavtive

}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(tags []SecurityTag) {
	for i := 0; i < len(tags); i++ {
		/**
		lessons Learnt: if you range over a slice/array, you're doing operation on the copy of the slice.
		*/
		deActivateTag(&tags[i].tag)
	}
}

func main() {

	//* Perform the following:
	//  - Create at least 4 items, all with active security tags

	tag1 := SecurityTag{name: "Tag 1", tag: Active}
	tag2 := SecurityTag{name: "Tag 2", tag: Active}
	tag3 := SecurityTag{name: "Tag 3", tag: Active}
	tag4 := SecurityTag{name: "Tag 4", tag: Active}

	//  - Store them in a slice or array
	tags := []SecurityTag{tag1, tag2, tag3, tag4}
	fmt.Println("All Tags:")
	fmt.Println(tags)

	//  - Deactivate any one security tag in the array/slice
	fmt.Println("Deactivated Tags[0]:")
	deActivateTag(&tags[0].tag)
	fmt.Println(tags)

	//  - Call the checkout() function to deactivate all tags
	fmt.Println("Checking Out Tags:")
	checkout(tags)
	fmt.Println(tags)
}
