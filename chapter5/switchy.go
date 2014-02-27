package main

import "fmt"

func main() {
	var input int
	fmt.Println("Tell me a number man? ") // User input!
	fmt.Scanf("%d", &input)               // %d for digit
	switch input {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
