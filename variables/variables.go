package main

func main() {
	// var name of vriable and its type then its value
	var name string = "Golang"
	println(name)

	// default value
	var isCool bool // default value is false
	println(isCool)

	// Default value of int is 0 and string is ""
	var score int
	println(score)
	println(name == "") // false
	println(score == 0) // true
	// Short declaration
	age := 10 // we use this when we have to declare and use its value at the same time
	println(age)

}
