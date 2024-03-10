package main

import (
	"fmt"
)

func main() {
	// var student1 string = "John" // string type
	// var student2 = "Jane"        // inferred
	// x := 2                       //inferred

	// var a string // default "" / blank
	// var b int    // default 0
	// var c bool   // default false

	// Assign after declare
	// var student3 string
	// student3 = "John"

	// Multivariable
	// var a, b, c, d int = 1, 3, 5, 6

	// Multivariable with different type
	// var a, b = 1, "hello"
	// c, d := 2, "world"

	// Declaration on Block
	var (
		a int
		b int    = 1
		c string = "hello"
		d string = "world"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
