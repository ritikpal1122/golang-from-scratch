package main

import "fmt"

// maps is associative data structure like dictionary in python

// map[keyType]valueType
func main() {

	// IMP : if key value is not present it will return zero value of that type
	// for string it will return "" and for int it will return 0 and for bool it will return false

	checkvalue := make(map[string]int)
	// declaring a map
	user := make(map[string]string)
	user["name"] = "Alice"

	fmt.Println(user["name"]) // Alice

	// creating a object
	person := map[string]string{
		"name":    "John",
		"country": "USA",
		"city":    "New York",
	}
	println(person["name"]) // John

	//deleting a key value pair
	delete(person, "city")
	// for example if key is not present it will return zero value of that type
	println(checkvalue["age"]) // 0

	// to clear a map we use clear function
	// but go does not have clear function so we have to reinitialize the map
	person = make(map[string]string)

	println(person["city"]) // ""

	// optional second values

	// to check if a key is present in a map we use the following syntax
	// this

	// we use _ if we dont want to use the value of the key
	value, ok := user["s"] // here value is the value of the key and ok is a boolean value that tells us if the key is present or not
	if ok {
		println("value is present", value)
	} else {
		println("value is not present")
	}
}
