package main

// range in go we use basically Iterating over data structures like arrays, slices, maps, strings, and channels.

func main() {
	// Example 1: Iterating over a slice
	nums := []int{1, 2, 3, 4, 5}

	// first for simple Loop
	for i := 0; i < len(nums); i++ {
		println("Index:", i, "Value:", nums[i])
	}

	// here index is the index of the element and value is the value of the element
	// Range returns two values: the index and the value of the element at that index.
	for index, value := range nums {
		println("Index:", index, "Value:", value)
	}

	// sum of elements in a slice
	sum := 0
	for _, value := range nums { // here we dont want the index so we use _
		sum += value
	}
	println("Sum:", sum)

	// Example 2: Iterating over a map
	user := map[string]string{
		"name":    "Alice",
		"country": "Wonderland",
		"city":    "Hearts",
	}
	for key, value := range user {
		println("Key:", key, "Value:", value)
	}
	// if we wnt only key then we can use _ for value
	for key := range user {
		println("Key:", key)
	}

	// Example 3: Iterating over a string
	// here i is the index of the character and c is the rune (unicode code point) of the character at that index
	for i, c := range "hello" {
		// here i is index and c is the character
		println(i, c)
	}

	// Example 4: Iterating over a channel
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // close the channel to avoid deadlock
	for value := range ch {
		println("Value from channel:", value)
	}
}
