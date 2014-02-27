package main

import "fmt"

func main() {
	// Bad naming below
	x := "Gary"
	fmt.Println(x, "the snail from Spongebob.")
	// Better naming
	name := "Gary"
	fmt.Println(name, "the snail from Spongebob.")
	// Even better naming
	snailName := "Gary"
	fmt.Println(snailName, "the snail from Spongebob.")
}
