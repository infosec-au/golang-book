package main

import "fmt"

func main() {
	// Go can infer the variable type and hence we don't need to spend time doing it manually
	// Suppose the longer form below, declaring x as a string assigned as "hello"

	var x string = "hello" // Doesn't this seem a bit too long?
	fmt.Println(x)
	// Let's shorten that down.
	y := "hello"   // No var type has been set! the : indicates that this is a new variable
	fmt.Println(y) // Same output man
	// The same goes for numbers
	z := 905
	fmt.Println(z)
	// And for every other variable type (e.g. floats)
	a := 0.12194
	fmt.Println(a)
}
