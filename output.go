package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"
	// var k, l int = 10, 20

	// "\n" create new line
	// fmt.Print(i, "\n")
	// fmt.Print(j, "\n")

	// Printing multiple variable on one print
	// fmt.Print(i, "\n", j)

	// add space between variable
	// fmt.Print(i, " ", j)

	// automatically add space when 2 variable are not string
	// fmt.Print(k, l)

	// Println() add whitespace and add newline
	// fmt.Println(i, j)

	// Printf() | Print Function
	// %v for value, %T for type
	fmt.Printf("i value : %v, j value %T", i, j)
}
