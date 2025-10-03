package main

func main() {
	// if else statement
	if true {
		println("This is true")
	} else {
		println("This is false")
	}

	age := 78
	if age < 18 {
		println("You are a minor")
	} else if age >= 18 && age < 60 {
		println("You are an adult")
	} else {
		println("You are a senior citizen")
	}
}
