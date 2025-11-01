package main

import "fmt"

// go have 5 type of reference types
//  1 Slices 2 Maps 3 Channels 4 Functions 5 Pointers

// we will go with all type of examples one by one
func main() {
	//1 pointers
	// pointers are used to store the address of a variable

	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	q := &j         // point to j
	*p = *q         // set i through the pointer
	fmt.Println(i)  // see the new value of i

}
