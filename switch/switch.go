package main

import "time"

func main() {

	// Simple switch case
	i := 5
	switch i {
	case 1:
		println("One")
	case 2:
		println("Two")
	case 3:
		println("Three")
	default:
		println("Other")
	}

	// Switch with multiple values

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend")
	default:
		println("It's a weekday Work Karo!")
	}

	// Switch with every weekday
	switch time.Now().Weekday() {
	case time.Monday:
		println("It's Monday")
	case time.Tuesday:
		println("It's Tuesday")
	case time.Wednesday:
		println("It's Wednesday")
	case time.Thursday:
		println("It's Thursday")
	case time.Friday:
		println("It's Friday")
	case time.Saturday:
		println("It's Saturday")
	case time.Sunday:
		println("It's Sunday")
	default:
		println("Unknown day")
	}

	// switch on odd even month
	switch time.Now().Month() {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		// print month name here as well

		println("It's an odd month", time.Now().Month().String())
	case time.April, time.June, time.September, time.November:
		println("It's an even month")
	case time.February:
		println("It's February")
	default:
		println("Unknown month")
	}

	// type switch

	WhoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("I am an integer")
		case string:
			println("I am a string")
		case bool:
			println("I am a boolean")
		default:
			println("I am of unknown type")
		}
	}
	WhoAmI(3)

}
