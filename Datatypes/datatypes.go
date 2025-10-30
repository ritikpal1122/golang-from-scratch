package main

import (
	"fmt"
)

func main() {
	var name string = "Golang"
	fmt.Println(name)

	var isCool bool
	fmt.Println(isCool)

	var score int
	fmt.Println(score)
	fmt.Println(name == "")
	fmt.Println(score == 0)

	age := 10
	fmt.Println(age)

	// Type of variables

	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of isCool is %T\n", isCool)

	// printf means print format
	// print ln means print line

	// aggregate types
	// Array Struct
	// Array
	var fruits [3]string = [3]string{"Apple", "Banana", "Grapes"}
	fmt.Println(fruits)
	// Struct
	type Person struct {
		name string
		age  int
	}
	var person1 Person = Person{name: "Ritik", age: 20}
	fmt.Println(person1)

	// Reference types
	// Slices Maps
	// Slices
	var numbers []int = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	// Maps
	var personMap map[string]int = map[string]int{"Ritik": 20, "Pal": 22}
	fmt.Println(personMap)

	// %v is used to print the value of any variable
	// %q is used to print the value of string with double quotes

	// Type conversion
	// in go everything is explicit no implicit type conversion

	var x int = 10
	var y float64 = float64(x)
	fmt.Println(y)

}
