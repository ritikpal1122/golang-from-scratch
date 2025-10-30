package main

import (
	"fmt"
	"slices"
)

// slices are like arrays but they are more flexible and powerful
// slices in go are dynamic arrays
// mostly used construct in go + usefukl methods

// A construct in go is a data structure that is used to store a collection of elements

func main() {
	// unintialized slice
	var nums []int       // declare a slice of integers
	println(nums == nil) // length of slice is 0

	// slice with values
	nums = []int{1, 2, 3, 4, 5}
	println(len(nums)) // length of slice is 5
	println(nums[0])   // first element of slice is 1

	// parameterized slice
	// type , length and capacity if we not provide capacity it will be same as length
	var client = make([]string, 3, 4) // make is a built-in function that creates a slice
	println(len(client))              // length of slice is 3

	client = append(client, "John", "Ritik", "3ed") // append is a built-in function that adds an element to the end of a slice
	// print slice values

	for i := range client {
		// dont print in new line
		print(client[i])
	}

	// slice is dynamic if we initialize its capacity and length both then it will be fixed size
	// but if we initialize only length then it will be dynamic
	// in the slice length can be less than or equal to capacity but not greater than capacity

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// copy lenght of another slice
	s = append(s, "d")

	copyvalue := make([]string, len((client)))
	copy(copyvalue, client)

	fmt.Println("copy:", copyvalue)

	// slice operator
	l := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("sl1:", l)      // entire slice
	fmt.Println("sl2:", l[2:5]) // from index 2 to 4
	fmt.Println("sl3:", l[:5])  // from start to index 4

	// slice compare
	t := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "s"}

	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	} else {
		fmt.Println("t and t2 are not equal")
	}

	// slice package (new in go 1.21
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	fmt.Println(slices.Equal(slice1, slice2)) // true

	// 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)
}
