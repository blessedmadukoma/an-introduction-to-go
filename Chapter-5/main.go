package main

import "fmt"

func main() {
	fmt.Println("Chapter 5 - Control Structures")
	// Print all numbers between 1 and 100 divisible by 3
	fmt.Println("\nMultiples of 3")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// FizzBuzz
	fmt.Println("\nFizzBuzz Solution")
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
