package main

//  Only construct in Go to perform iteration is the for loop. Or Perform Looping in Go

// executing a set of statements multiple times until a certain condition is met

func main() {

	// this is while loop in other languages
	i := 1
	for i <= 3 {
		println(i)
		i++
	}

	// Infinite Loop
	for {
		print("Hello, ")
		break
	}

	// 2nd infinite loop example
	sum := 1
	for {
		if (sum) >= 10 {
			break
		}
		sum += sum
	}
	println("this si si", sum)
	// Classic For Loop
	for i := 0; i < 5; i++ {
		println(i)

	}

	// 1.22  New feature Range that basically iterates over a slice or map
	for i := range 10 { // It will print 0 to 9 because range starts from 0
		println(i)
	}

}
