package main

import "fmt"

// functions in go are defined using the func keyword
//  go requires every program to have a main function
// go required explicit return type if we dont provide it will throw an error

// Structure of the funtion
// func name of fun and then name of parameter and type w

func plus(a int, b int) int { // here 3 int first two are parameters and last one is return type
	return a + b
}

func swaps(x, y string) (string, string) {
	return y, x
}

func vals() (int, int) {
	return 3, 7
}

// a fun that return multiple values // we do not use this format often
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// for that we use explicit return

func explicitReturn(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

func main() {

	res := plus(3, 4) // here 3 and 4 are arguments
	println(res)

	// return multiple values in go
	x, y := swaps("hello", "world")
	println(x, y)

	a, b := vals()
	println(a, b)

	fmt.Println(split(17))
}
