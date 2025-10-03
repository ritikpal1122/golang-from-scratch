package main

// constants are immutable values which are known at compile time and do not change for the life of the program
func main() {
	const pi = 3.14
	println(pi)
	const language string = "Golang"
	println(language)

	// language ="Java" // error: cannot assign to language (constant variable)

	// for multiple constants we use constant grouping
	const (
		a = 1
		b = 2
	)
	println(a, b)

}
