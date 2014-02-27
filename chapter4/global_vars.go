package main

import "fmt"

var x string = "This is a globally accessable variable string!" // This is accessable by all functions

const y string = "This is a globally accessable constant string!" // This is accessable by all functions, but this string cannot change, it is *constant*

func main() {
	fmt.Println("OMG, I can access x:", x)
	fmt.Println("yey, let's print a const:", y)
}
func f() {
	fmt.Println("Lol, settle down! I can access x too:", x)
}
