package main

import (
	"fmt"
)

func main() {
	// Question 2
	fmt.Println("The Largest 8 digit number in binary:")
	x := (10 ^ 1)
	x = x - 1
	fmt.Println(x)

	// Question 3
	fmt.Println(321325 * 424521)

	// Question 4: Finding the length of a string
	fmt.Println(len("Blessed Madukoma"))

	// Question 5
	fmt.Println((true && false) || (false && true) || !(false && false)) // true
}
