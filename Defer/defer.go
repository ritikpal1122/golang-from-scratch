package main

func main() {

	defer println("World") // this will be executed at the end of the function

	// pushing the defer statements onto a stack after this they will be executed in LIFO order
	println("Hello")
}
