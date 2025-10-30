package main

import "fmt"

var c, python, java bool

func main() {
	var i int

	j := 0                             // this is only accesable inside main function cannot be used outside main
	fmt.Println(i, c, python, java, j) // prints 0
}
