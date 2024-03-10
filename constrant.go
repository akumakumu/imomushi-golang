package main

import (
	"fmt"
)

// Untyped
// const PI = 3.14

// Typed
// const A int = 1

// Multiple Declaration
const (
	A int = 1
	B     = 3.14
	C     = "Hi"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
