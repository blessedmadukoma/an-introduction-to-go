package main

import "fmt"

func main() {
	fmt.Println("Chapter 4 - Variables")
	x := 5
	x += 1
	fmt.Println(x) // x = 6

	// Converting Fahrenheit into Celsius
	var C, F float64
	F = 33
	C = (F - 32) * 5 / 9
	fmt.Println(C)

	// Converting feet to meters
	var ft, val float64
	m := 0.3048
	ft = 2
	val = ft * m
	fmt.Println(val)
}
