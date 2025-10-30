package main

import "fmt"

// array in go basically number sequence of elements of same type
func main() {

	var nums [5]int
	nums[0] = 1          // array of 5 integers
	fmt.Println(nums[0]) // length of array is 5

	num := [5]int{1, 2, 3, 4, 5} // array of 5 integers
	fmt.Println(num)

	// 2 D array
	var matrix [2][3]int
	fmt.Println(matrix)
	matrix[0][0] = 1
	fmt.Println(matrix)
	matrix[0][1] = 2
	fmt.Println(matrix)
	matrix[0][2] = 3
	fmt.Println(matrix)
	matrix[1][0] = 4
	fmt.Println(matrix)
}
